/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package daos

import (
	"errors"
	"fmt"
	"time"

	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/databases"
	"github.com/huavanthong/microservice-golang/user-api/models"
	"github.com/huavanthong/microservice-golang/user-api/security"
	"github.com/huavanthong/microservice-golang/user-api/utils"
	"gopkg.in/mgo.v2/bson"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {
	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// query all users in MongoDB and store it to array
	var users []models.User
	err := collection.Find(bson.M{}).All(&users)

	return users, err
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (models.User, error) {

	// validate user id
	err := u.utils.ValidateObjectID(id)
	if err != nil {
		return models.User{}, err
	}

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// find a user by id
	var user models.User
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)

	return user, err
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {

	// validate user id
	err := u.utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// delete user by id
	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	return err
}

// Login User
func (u *User) Login(ac models.Account) (models.User, error) {

	var err error
	var user models.User

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	/********* Design 1: Get user info with username and password *********/
	if ac.UserName == "admin" {
		err = collection.Find(bson.M{"$and": []bson.M{
			bson.M{"name": ac.UserName},
			bson.M{"email": ac.Email},
			bson.M{"password": ac.Password}},
		}).One(&user)

	} else {
		/********* Design 2: Get user info only with username, then check password by bcrypt *********/
		var result bson.M
		err = collection.Find(bson.M{"name": ac.UserName}).One(&result)
		fmt.Println("Check result: ", result)
		if result == nil {
			return user, errors.New("User not found")

		} else {
			// convert interface to string
			hashedPassword := fmt.Sprintf("%v", result["password"])

			err = security.CheckPasswordHash(hashedPassword, ac.Password)
			if err != nil {
				return user, err
			}
		}
	}

	return user, err
}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {

	// define error code
	var err error = nil

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()

	// get the User collection to execute the query against.
	collectionUser := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// check if a user already exists
	var uesrExist models.User
	err = collectionUser.Find(bson.M{"$and": []bson.M{
		bson.M{"name": user.Name},
		bson.M{"email": user.Email}}}).One(&uesrExist)
	if err == nil {
		return errors.New("Email or Username already exists")
	}

	// hash password using bcrypt
	password, serr := security.Hash(user.Password)
	if serr != nil {
		return errors.New(common.ErrHashPasswordFail)
	}

	// update basic info for new user
	var newUser models.User = user
	newUser.Password = password
	newUser.Role = models.Role{
		RoleName: "member",
		Actions: []models.Action{
			{
				ActionName: "read, write",
			},
		},
	}
	newUser.CreatedAt = time.Now().UTC().String()

	// insert a new user from argument
	err = collectionUser.Insert(&newUser)
	if err != nil {
		return errors.New("Insert new user failed")
	} else {

		// get a Profile collection to execute the query against.
		collectionProfile := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

		// create a new profile from basic info user
		profile := models.Profile{
			ID:        bson.NewObjectId(),
			FirstName: newUser.Name,
			Email:     newUser.Email,
			UserID:    newUser.ID,
		}

		// insert a profile from user info
		err = collectionProfile.Insert(&profile)
		if err != nil {
			return errors.New("Insert profile for new user failed")
		}
	}

	return err

}

// Delete remove an existing User
func (u *User) Delete(user models.User) error {
	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// remove a user match entired info
	err := collection.Remove(&user)
	return err
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// update user by id
	err := collection.UpdateId(user.ID, &user)
	return err
}

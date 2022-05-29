/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package daos

import (
	"errors"
	"fmt"

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
func (u *User) Login(name string, password string) (models.User, error) {

	var err error
	var user models.User

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	/********* Design 1: Get user info with username and password *********/
	if name == "admin" {
		err = collection.Find(bson.M{"$and": []bson.M{
			bson.M{"name": name},
			bson.M{"password": password}},
		}).One(&user)

	} else {
		/********* Design 2: Get user info only with username, then check password by bcrypt *********/
		var result bson.M
		err = collection.Find(bson.M{"name": name}).One(&result)
		fmt.Println(err, result)
		fmt.Println(result["name"])

		// convert interface to string
		hashedPassword := fmt.Sprintf("%v", result["password"])

		err = security.CheckPasswordHash(hashedPassword, password)
		if err != nil {
			return user, err
		}
	}

	return user, err
}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {

	// hash password using bcrypt
	password, serr := security.Hash(user.Password)
	if serr != nil {
		return errors.New(common.ErrHashPasswordFail)
	}

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// update data for new user
	var newUser models.User = user
	newUser.Password = password

	// insert a new user from argument
	err := collection.Insert(&newUser)
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

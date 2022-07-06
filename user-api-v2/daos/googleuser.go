/*
 * @File: daos.googleuser.go
 * @Description: Implements User Social from Google for MongoDB
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package daos

import (
	"fmt"

	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/databases"
	"github.com/huavanthong/microservice-golang/user-api/models"
	"github.com/huavanthong/microservice-golang/user-api/utils"
	"gopkg.in/mgo.v2/bson"
)

// User Info from Google
type GoogleUser struct {
	utils *utils.Utils
}

// SaveUser register a user so we know that we saw that user already.
func (gu *GoogleUser) SaveUser(u *models.GoogleUser) error {

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// check the existed user
	if _, err := gu.LoadUser(u.Email); err == nil {
		return fmt.Errorf("user already exists")
	}

	// convert google user info to server db
	user := models.User{
		ID:    bson.NewObjectId(),
		Name:  u.Name,
		Email: u.Email,
	}

	profile := models.Profile{
		ID:            bson.NewObjectId(),
		FirstName:     user.Name,
		UserID:        user.ID,
		Email:         user.Email,
		EmailVerified: u.EmailVerified,
		Picture:       u.Picture,
		Gender:        u.Gender,
	}

	// get a collection to execute the query against.
	userCollection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)
	err := userCollection.Insert(user)

	// if not error, update to profile
	if err == nil {
		profileCollection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)
		err = profileCollection.Insert(profile)
	}

	return err
}

// LoadUser get data from a user.
func (gu *GoogleUser) LoadUser(Email string) (u *models.GoogleUser, err error) {

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	c := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err = c.Find(bson.M{"email": Email}).One(&u)

	return u, err
}

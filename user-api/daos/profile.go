/*
 * @File: daos.profile.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package daos

import (
	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/databases"
	"github.com/huavanthong/microservice-golang/user-api/models"
	"github.com/huavanthong/microservice-golang/user-api/utils"
	"gopkg.in/mgo.v2/bson"
)

// Profile manages User CRUD
type Profile struct {
	utils *utils.Utils
}

// GetProfileByUserId gets the profile by specific user id
func (p *Profile) GetProfileByUserId() (models.Profile, error) {
	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

	// query all users in MongoDB and store it to array
	var profile models.Profile
	err := collection.Find(bson.M{}).All(&profile)

	return profile, err
}

// Insert adds a new profile by specific user into database'
func (u *Profile) Insert(profile models.Profile) error {

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

	// insert a profile info to specific user id
	err := collection.Insert(&profile)
	return err

}

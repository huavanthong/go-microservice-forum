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
func (p *Profile) GetProfileByUserId(userid string) (models.Profile, error) {

	// validate user id
	err := p.utils.ValidateObjectID(userid)
	if err != nil {
		return models.Profile{}, err
	}

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// find a user by id
	var profile models.Profile
	err = collection.FindId(bson.ObjectIdHex(userid)).One(&profile)

	return profile, err
}

// Insert adds a new profile by specific user into database'
func (p *Profile) Insert(profile models.Profile) error {

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

	// insert a profile info to specific user id
	err := collection.Insert(&profile)
	return err

}

// DeleteByID finds a Profile by user id
func (p *Profile) DeleteByUserID(id string) error {

	// validate user id
	err := p.utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

	// delete user by id
	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	return err
}

// Update modifies an existing User
func (p *Profile) Update(profile models.Profile) error {
	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

	// update user by id
	err := collection.UpdateId(profile.ID, &profile)
	return err
}

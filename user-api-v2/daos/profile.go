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

// GetAll gets the list of Profiles
func (u *Profile) GetAll() ([]models.Profile, error) {
	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

	// query all users in MongoDB and store it to array
	var profiles []models.Profile
	err := collection.Find(bson.M{}).All(&profiles)

	return profiles, err
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

	// get the Profile collection to execute the query against
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

	// find the profile user by userid
	var profile models.Profile
	err = collection.Find(bson.M{"$and": []bson.M{
		bson.M{"_userid": bson.ObjectIdHex(userid)}},
	}).One(&profile)

	return profile, err
}

// Update modifies an existing User
func (p *Profile) Update(profile models.Profile) error {
	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColProfile)

	// update profile user by userid
	err := collection.UpdateId(bson.M{"_userid": profile.UserID}, &profile)

	return err
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

/*
 * @File: daos.googleuser.go
 * @Description: Implements User Social from Google for MongoDB
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package daos

import (
	"fmt"

	"github.com/huavanthong/microservice-golang/user-api/databases"
	"github.com/huavanthong/microservice-golang/user-api/models"
	"github.com/huavanthong/microservice-golang/user-api/utils"
	"gopkg.in/mgo.v2/bson"
)

// User manages User CRUD
type GoogleUser struct {
	utils *utils.Utils
}

// SaveUser register a user so we know that we saw that user already.
func (gu *GoogleUser) SaveUser(u *models.GoogleUser) error {

	// copy for a newsession with original authentication
	// to access to MongoDB.
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	if _, err := mdb.LoadUser(u.Email); err == nil {
		return fmt.Errorf("user already exists")
	}
	c := mdb.session.DB("webadventure").C("users")
	err := c.Insert(u)
	return err
}

// LoadUser get data from a user.
func (mdb MongoDBConnection) LoadUser(Email string) (result structs.User, err error) {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	c := mdb.session.DB("webadventure").C("users")
	err = c.Find(bson.M{"email": Email}).One(&result)
	return result, err
}

// GetSession return a new session if there is no previous one.
func (mdb *MongoDBConnection) GetSession() *mgo.Session {
	if mdb.session != nil {
		return mdb.session.Copy()
	}
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

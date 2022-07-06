/*
 * @File: databases.mongodb.go
 * @Description: Handles MongoDB connections
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package databases

import (
	"time"

	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/models"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbSession  *mgo.Session
	Databasename string
}

// Init initalizes mongo database
func (db *MongoDB) Init() error {
	db.Databasename = common.Config.MgDbName

	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{common.Config.MgAddrs}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: db.Databasename,            // Database name
		Username: common.Config.MgDbUsername, // Username
		Password: common.Config.MgDbPassword, // Password
	}

	// Create a session which maintains a pool of socket connections
	// to the DB MongoDB database.
	var err error
	db.MgDbSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Debug("Can't connect to mongo, go error: ", err)
		return err
	}

	return db.initData()
}

// InitData initializes default data
func (db *MongoDB) initData() error {
	var err error
	var count int

	// Check if user collection has at least one document
	sessionCopy := db.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Databasename).C(common.ColUsers)
	count, err = collection.Find(bson.M{}).Count()

	if count < 1 {
		// Create admin/admin account
		user := models.User{
			ID:            bson.NewObjectId(),
			Name:          "admin",
			Email:         "admin@gmail.com",
			Password:      "admin",
			LoginAttempts: []models.LoginAttempt{}, // avoid struct composite literals. More details:https://stackoverflow.com/questions/51742263/go-build-with-protocol-buffer-error-too-few-values-in-struct-initializer

		}
		err = collection.Insert(&user)
	}

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Close()
	}
}

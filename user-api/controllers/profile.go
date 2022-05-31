/*
 * @File: controllers.profile.go
 * @Description: Implements User API logic functions
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/models"
	"github.com/huavanthong/microservice-golang/user-api/payload"
	"github.com/huavanthong/microservice-golang/user-api/utils"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

// Define profile manages
type Profile struct {
	utils       utils.Utils
	profileDaos daos.Profile
}

func (u *Profile) AddProfile(ctx *gin.Context) {
	// bind profile info to json getting context
	var p models.Profile
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	// validate data on profile of user
	v := utils.NewValidation()

	err := v.Validate(p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	address := &models.Address{
		Street:   "Kinh Duong Vuong",
		Ward:     "12",
		District: "6",
		City:     "Ho Chi Minh",
		Country:  "Viet Nam",
	}

	// create profile from models
	profile := models.Profile{
		bson.NewObjectId(),
		p.ProfileName,
		p.FirstName,
		p.LastName,
		p.Email,
		p.AccountID,
		p.Age,
		p.PhoneNumber,
		p.DefaultProfile,
		p.FavouriteColor,
		address}

	// insert user to DB
	err = u.profileDaos.Insert(profile)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, payload.Message{"Successfully"})
		log.Debug("Create profile = " + profile.Name + ", password = " + profile.Password)
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

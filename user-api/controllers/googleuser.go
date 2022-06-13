package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/daos"
	"github.com/huavanthong/microservice-golang/user-api/models"
	"github.com/huavanthong/microservice-golang/user-api/payload"

	googleCred "github.com/huavanthong/microservice-golang/user-api/security/google"
	"github.com/huavanthong/microservice-golang/user-api/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Define user manages
type GoogleUser struct {
	utils    utils.Utils
	guserDAO daos.GoogleUser
}

/**************************************************************************************
 * global configuration for google security.
/*************************************************************************************/
var cred googleCred.Credentials
var conf *oauth2.Config

/**************************************************************************************
 * Internal function
/*************************************************************************************/
func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func init() {
	file, err := ioutil.ReadFile("./config/google-credentials.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &cred); err != nil {
		log.Println("unable to marshal data")
		return
	}

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:8808/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

/**************************************************************************************
 * RESTful API
/*************************************************************************************/
// IndexHandler handles the location /.
func (gu *GoogleUser) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

// AuthGoogleAccount godoc
// @Summary Check user authentication
// @Description Authenticate google user
// @Tags admin
// @Security ApiKeyAuth
// @Accept  json
// @Failure 401 {object} payload.Error
// @Failure 500 {object} payload.Error
// @Success 200 {object} security.Token
// @Router /admin/auth [get]
// AuthGoogleAccount handles authentication of a user and initiates a session.
func (gu *GoogleUser) AuthGoogleAccount(ctx *gin.Context) {
	// Handle the exchange code to initiate a transport.
	// get session where stored info users
	session := sessions.Default(ctx)

	// retrieve token for accessing google service
	retrievedState := session.Get("state")
	queryState := ctx.Request.URL.Query().Get("state")

	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		ctx.JSON(http.StatusUnauthorized, payload.Error{common.StatusCodeUnknown, common.InvalidSession})
		// ctx.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Invalid session state."})
		return
	}
	code := ctx.Request.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, common.LoginFailed})
		// ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Login failed. Please try again."})
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	u := models.GoogleUser{}
	if err = json.Unmarshal(data, &u); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, common.MarshallingFailed})
		// ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error marshalling response. Please try agian."})
		return
	}
	session.Set("user-id", u.Email)
	err = session.Save()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, common.SavingSessionError})
		// ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving session. Please try again."})
		return
	}
	seen := false
	if _, mongoErr := gu.guserDAO.LoadUser(u.Email); mongoErr == nil {
		seen = true
	} else {
		err = gu.guserDAO.SaveUser(&u)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, common.SavingUserError})
			// ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving user. Please try again."})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"email": u.Email, "seen": seen})
	// ctx.HTML(http.StatusOK, "battle.tmpl", gin.H{"email": u.Email, "seen": seen})
}

// LoginGoogle godoc
// @Summary Check token for accessing google account
// @Description get token for redirect to sign in google service
// @Tags admin
// @Security ApiKeyAuth
// @Accept  json
// @Failure 401 {object} payload.Error
// @Failure 500 {object} payload.Error
// @Success 200 {object} security.Token
// @Router /admin/auth/social [get]
// LoginGoogle handles the login procedure.
func (gu *GoogleUser) LoginGoogle(ctx *gin.Context) {

	// State is a token to protect the user from CSRF attachks
	// Refer: https://pkg.go.dev/golang.org/x/oauth2#section-readme
	state, err := googleCred.RandToken(32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error while generating random data."})
		// ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Error while generating random data."})
		return
	}

	// Getting session default from main.go: NewCookieStore()
	session := sessions.Default(ctx)

	// Assign token to keyword: state
	session.Set("state", state)

	// Save session
	err = session.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	// get URL to redirect sign up google
	link := getLoginURL(state)
	ctx.JSON(http.StatusOK, gin.H{"link": link})
	// ctx.HTML(http.StatusOK, "auth.tmpl", gin.H{"link": link})
}

// FieldHandler is a rudementary handler for logged in users.
func (gu *GoogleUser) FieldHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user-id")
	ctx.HTML(http.StatusOK, "field.tmpl", gin.H{"user": userID})
}

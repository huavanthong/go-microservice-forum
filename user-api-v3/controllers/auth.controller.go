package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"

	"github.com/huavanthong/microservice-golang/user-api-v3/config"
	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"github.com/huavanthong/microservice-golang/user-api-v3/payload"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
	"github.com/huavanthong/microservice-golang/user-api-v3/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
	ctx         context.Context
	collection  *mongo.Collection
}

func NewAuthController(authService services.AuthService, userService services.UserService, ctx context.Context, collection *mongo.Collection) AuthController {
	return AuthController{authService, userService, ctx, collection}
}

// SignUpUser godoc
// @Summary Register a new user
// @Description Register a new user for service
// @Tags auth
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param user body models.SignUpInput true "New User"
// @Failure 400 {object} payload.Response
// @Failure 409 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 201 {object} payload.UserRegisterSuccess
// @Router /auth/register [post]
// SignUp User
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var user *models.SignUpInput

	// from context, bind user info to json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// confirm password
	if user.Password != user.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: "Passwords do not match",
			})
		return
	}

	// transfer user info to service
	newUser, err := ac.authService.SignUpUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusConflict,
					Message: err.Error(),
				})
			return
		}

		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}
	/********************** Verify email *********************/
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	// Generate Verification Code
	code := randstr.String(20)

	verificationCode := utils.Encode(code)

	// Update User in Database
	ac.userService.UpdateUserById(newUser.ID.Hex(), "verificationCode", verificationCode)

	var firstName = newUser.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// 👇 Send Email
	emailData := utils.EmailData{
		URL:       config.Origin + "/api/v3/auth/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	utils.SendEmail(newUser, &emailData)

	message := "We sent an email with a verification code to " + user.Email

	// return user info after register a new user successfully
	ctx.JSON(http.StatusCreated,
		payload.UserRegisterSuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: message,
			Data:    models.FilteredResponse(newUser),
		})
}

// SignInUser godoc
// @Summary Sign In User
// @Description User sign in to service
// @Tags auth
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param user body models.SignInInput true "Authenticate user"
// @Failure 400 {object} payload.Response
// @Failure 401 {object} payload.Response
// @Success 200 {object} payload.UserLoginSuccess
// @Router /auth/login [post]
// SignIn User
func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var credentials *models.SignInInput

	// from context, bind user info to json
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// Find user by email
	user, err := ac.userService.FindUserByEmail(credentials.Email)
	fmt.Println(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusBadRequest,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusBadRequest,
					Message: "Invalid email or password",
				})
			return
		}
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// User'email verify or not
	fmt.Println(user.Verified)
	if !user.Verified {
		ctx.JSON(http.StatusUnauthorized,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusUnauthorized,
				Message: "You are not verified, please verify your email to logi",
			})
		return
	}

	// If user exists, verify password
	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: "Invalid email or Password",
			})
		return
	}

	// loading config, getting private key for generating token
	config, _ := config.LoadConfig(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// set to cookie
	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK,
		payload.UserLoginSuccess{
			Status:      "success",
			Message:     "Generate token success",
			AccessToken: access_token,
		})
}

// RefreshAccessToken godoc
// @Summary Refresh access token
// @Description Refresh access token after the specific time
// @Tags auth
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Failure 403 {object} payload.Response
// @Success 200 {object} payload.UserRefreshTokenSuccess
// @Router /auth/refresh [get]
// Refresh Access Token
func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: message,
			})
		return
	}

	config, _ := config.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	user, err := ac.userService.FindUserById(fmt.Sprint(sub))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: "the user belonging to this token no logger exists",
			})
		return
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK,
		payload.UserRefreshTokenSuccess{
			Status:      "success",
			Message:     "Refresh token success",
			AccessToken: access_token,
		})
}

// GoogleOAuth godoc
// @Summary Sign in a new user by Google OAuth2
// @Description Sign in a new user by Google OAtuth2, then save a new user to DB
// @Tags auth
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Failure 307 {object} payload.Response
// @Failure 400 {object} payload.Response
// @Failure 401 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 201 {string} http.StatusCreated
// @Router /sessions/oauth/google [get]
// SignUp User by Google
func (ac *AuthController) GoogleOAuth(ctx *gin.Context) {
	code := ctx.Query("code")
	var pathUrl string = "/"

	if ctx.Query("state") != "" {
		pathUrl = ctx.Query("state")
	}

	if code == "" {
		ctx.JSON(http.StatusUnauthorized,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusUnauthorized,
				Message: "Authorization code not provided!",
			})
		return
	}

	// Use the code to get the id and access tokens
	tokenRes, err := utils.GetGoogleOauthToken(code)

	if err != nil {
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
	}

	user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
	}

	createdAt := time.Now()
	resBody := &models.UpdateDBUser{
		Email:     user.Email,
		Name:      user.Name,
		Photo:     user.Picture,
		Provider:  "google",
		Role:      "user",
		Verified:  true,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	updatedUser, err := ac.userService.UpsertUser(user.Email, resBody)
	if err != nil {
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
	}

	config, _ := config.LoadConfig(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, updatedUser.ID.Hex(), config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, updatedUser.ID.Hex(), config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprint(config.ClientOrigin, pathUrl))
}

// LogoutUser godoc
// @Summary Log out user
// @Description Delete all cookie in session
// @Tags auth
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 201 {string} StatusOK
// @Router /logout [get]
func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// VerifyEmail godoc
// @Summary Verify email user
// @Description Verify email user that sign up to service
// @Tags auth
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param verificationCode path string true "Verification Code"
// @Failure 403 {object} payload.Response
// @Success 209 {object} payload.Response
// @Router /verifyemail/{verificationCode} [get]
func (ac *AuthController) VerifyEmail(ctx *gin.Context) {

	code := ctx.Params.ByName("verificationCode")
	verificationCode := utils.Encode(code)

	query := bson.D{{Key: "verificationCode", Value: verificationCode}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "verified", Value: true}}}, {Key: "$unset", Value: bson.D{{Key: "verificationCode", Value: ""}}}}
	result, err := ac.collection.UpdateOne(ac.ctx, query, update)
	if err != nil {
		ctx.JSON(http.StatusForbidden,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusForbidden,
				Message: err.Error(),
			})
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusForbidden,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusForbidden,
				Message: "Could not verify email address",
			})
		return
	}

	fmt.Println(result)

	ctx.JSON(http.StatusOK,
		payload.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Email verified successfully",
		})
}

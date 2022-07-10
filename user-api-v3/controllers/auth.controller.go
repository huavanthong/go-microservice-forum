package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api-v3/config"
	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"github.com/huavanthong/microservice-golang/user-api-v3/payload"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
	"github.com/huavanthong/microservice-golang/user-api-v3/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(authService services.AuthService, userService services.UserService) AuthController {
	return AuthController{authService, userService}
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
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	// return user info after register a new user successfully
	ctx.JSON(http.StatusCreated,
		payload.UserRegisterSuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: "Register a new user successfully",
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
// SignUp User
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

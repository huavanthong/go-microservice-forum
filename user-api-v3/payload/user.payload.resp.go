/*
 * @File: user.payload.resp.go
 * @Description: Return payload info for user
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

/****************************** Reference ****************************************
Refer: [here](https://cloud.google.com/storage/docs/json_api/v1/status-codes#401-unauthorized)
Error example: [here](https://github.com/googleapis/google-cloud-go/blob/main/cmd/go-cloud-debug-agent/internal/controller/client_test.go)
Swagger message example: [here](https://docs.swagger.io/spec.html#511-object-example)
******************************************************************************/

import "github.com/huavanthong/microservice-golang/user-api-v3/models"

// Error defines the response error
type UserRegisterSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Register a new user successfully"`
	Data    models.UserResponse
}

type UserLoginSuccess struct {
	Status      string `json:"status" example:"success"`
	Message     string `json:"message" example:"Generate token success"`
	AccessToken string `json:"access_token" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc0NTg2NzcsImlhdCI6MTY1NzQ1Nzc3NywibmJmIjoxNjU3NDU3Nzc3LCJzdWIiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAifQ.WbRHMAdggCfHR06XKpmbFCu3DNjPkjOPYs9b8TuvBZym1d_TD7JCMadmNCq_Sim9bOzhMi8muDZBb1wRBkli2A"`
}

type UserRefreshTokenSuccess struct {
	Status      string `json:"status" example:"success"`
	Message     string `json:"message" example:"Refresh token success"`
	AccessToken string `json:"access_token" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc0NTg2NzcsImlhdCI6MTY1NzQ1Nzc3NywibmJmIjoxNjU3NDU3Nzc3LCJzdWIiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAifQ.WbRHMAdggCfHR06XKpmbFCu3DNjPkjOPYs9b8TuvBZym1d_TD7JCMadmNCq_Sim9bOzhMi8muDZBb1wRBkli2A"`
}

type GetUserSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get user success"`
	Data    models.UserResponse
}

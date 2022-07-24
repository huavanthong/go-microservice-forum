/*
 * @File: admin.payload.resp.go
 * @Description: Return payload info for admin
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
type GetAllUsersSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get all users success"`
	Data    []models.User
}

type AdminGetUserSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get user success"`
	Data    models.User
}

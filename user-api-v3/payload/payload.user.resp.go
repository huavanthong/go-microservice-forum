/*
 * @File: payload.user.resp.go
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
type UserRegister struct {
	Data            models.UserResponse
	ResponseMessage Response
}

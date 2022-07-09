/*
 * @File: payload.responseMessage.go
 * @Description: Defines Error information will be returned to the clients
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

/****************************** Reference ****************************************
Refer: [here](https://cloud.google.com/storage/docs/json_api/v1/status-codes#401-unauthorized)
Error example: [here](https://github.com/googleapis/google-cloud-go/blob/main/cmd/go-cloud-debug-agent/internal/controller/client_test.go)
Swagger message example: [here](https://docs.swagger.io/spec.html#511-object-example)
******************************************************************************/

// Error defines the response error
type ResponseMessage struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"27"`
	Message string `json:"message" example:"Error message"`
	// Errors  []ErrorItem `json:"errors"`
}

type ErrorItem struct {
	Domain       string `json:"domain" example:"user-api-v3"`
	Reason       int    `json:"reason" example:"Failed due to ..."`
	Message      string `json:"message" example:"invalidParameter"`
	LocationType string `json:"locationtype" example:"region"`
	Location     string `json:"location" example:"SOUTHAMERICA-EAST1"`
}

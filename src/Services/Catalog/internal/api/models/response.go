/*
 * @File: payload.response.go
 * @Description: Defines Error information will be returned to the clients
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

/****************************** Reference ****************************************
Refer: [here](https://cloud.google.com/storage/docs/json_api/v1/status-codes#401-unauthorized)
Error example: [here](https://github.com/googleapis/google-cloud-go/blob/main/cmd/go-cloud-debug-agent/internal/controller/client_test.go)
Swagger message example: [here](https://docs.swagger.io/spec.html#511-object-example)
******************************************************************************/

// Error defines the response error
type Response struct {
	Status  string `json:"status" example:"failed"`
	Code    int    `json:"code" example:"400"`
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

type ResponseSuccess struct {
	Status  string      `json:"statusCode"`
	Code    int         `json:"method"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Responses struct {
	StatusCode int    `json:"statusCode"`
	Method     string `json:"method"`
	Message    string `json:"message"`
}

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

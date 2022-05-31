/*
 * @File: payload.message.go
 * @Description: Defines Message information will be returned to the clients
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

// Message defines the response message
type Message struct {
	Message string `json:"message" example:"message"`
}

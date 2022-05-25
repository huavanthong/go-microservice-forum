/*
 * @File: payload.error.go
 * @Description: Defines Error information will be returned to the clients
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

// Error defines the response error
type Error struct {
	Code    int    `json:"code" example:"27"`
	Message string `json:"message" example:"Error message"`
}

/*
 * @File: security.token.go
 * @Description: Defines Token information will be returned to the clients
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package security

// Token string
type Token struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicmF5Y2FkIiwicm9sZSI6IiIsImV4cCI6MTUzOTI0OTc3OSwiaXNzIjoic2VlZG90ZWNoIn0.lVHq8J_0vfvECbplogAOCDCNh63ivTBOhya8KE6Ew_E"`
}

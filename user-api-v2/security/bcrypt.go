/*
 * @File: security.bcrypt.go
 * @Description: Defines Token information will be returned to the clients
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */

package security

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Hash a password with bcrypt
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Client compare a submit password
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Avoid SQL Injection by using santize
func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

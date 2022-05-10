/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package daos

import (
	"../models"
)

// User manages User CRUD
type User struct {
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {

	return nil, nil
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (models.User, error) {

	return models.User{}, nil
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {

	return nil
}

// Login User
func (u *User) Login(name string, password string) (models.User, error) {

	return models.User{}, nil
}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {
	return nil
}

// Delete remove an existing User
func (u *User) Delete(user models.User) error {
	return nil
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	return nil
}

/*
 * @File: models.user_test.go
 * @Description: Test User model
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

import (
	"errors"
	"fmt"
	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/stretchr/testify/assert"

	"testing"
)

/*********************** Account ***********************/
func TestNormalCaseAddUser(t *testing.T) {

	a := Account{
		Name:     "hvthong",
		Password: "1234",
	}

	err := a.Validate()
	assert.NoError(t, err)
}

func TestInvalidUserNameReturnErr(t *testing.T) {

	a := Account{
		Name:     "",
		Password: "1234",
	}

	err := a.Validate()

	assert.Equal(t, err, errors.New(common.ErrNameEmpty))
}

func TestInvalidPassWordReturnErr(t *testing.T) {

	a := Account{
		Name:     "hvthong",
		Password: "",
	}

	err := a.Validate()

	assert.Equal(t, err, errors.New(common.ErrPasswordEmpty))
}

func TestSqlInjection(t *testing.T) {

	a := Account{
		Name:     "hvthong > ok",
		Password: "1234$123>33",
	}

	err := a.Validate()

	fmt.Println(a.Name)
	fmt.Println(a.Password)

	assert.NoError(t, err)
}

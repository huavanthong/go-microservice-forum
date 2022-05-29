/*
 * @File: models.profiles_test.go
 * @Description: Test Profiles model
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package daos

import (
	"fmt"
	"github.com/huavanthong/microservice-golang/user-api/utils"
	"testing"
)

/*********************** Normal Case: Create a profile ***********************/
func TestNormalCaseGetAllUsers(t *testing.T) {

	var utils utils.Utils

	user := &User{utils}
	user.GetAll()

	fmt.Println(user)
}

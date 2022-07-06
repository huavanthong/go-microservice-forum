/*
 * @File: models.profiles_test.go
 * @Description: Test Profiles model
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

import (
	"github.com/huavanthong/microservice-golang/user-api/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

/*********************** Normal Case: Create a profile ***********************/
func TestNormalCaseCreateProfile(t *testing.T) {

	address := &Address{
		Street:   "Kinh Duong Vuong",
		Ward:     "12",
		District: "6",
		City:     "Ho Chi Minh",
		Country:  "Viet Nam",
	}

	profile := &Profiles{
		ProfileID:      1,
		ProfileName:    "Personal Profile",
		FirstName:      "Thong",
		LastName:       "Hua Van",
		Email:          "Badger.Smith@gmail.com",
		AccountID:      1,
		Age:            30,
		PhoneNumber:    "0908354129",
		DefaultProfile: "Personal Profile",
		FavouriteColor: "#ff0000",
		Addresses:      []*Address{address},
	}

	v := utils.NewValidation()

	err := v.Validate(profile)

	assert.Len(t, err, 0)
}

/*********************** Adnormal Case: Create a profile ***********************/
func TestAbnormalCaseAgeGreater(t *testing.T) {

	address := &Address{
		Street:   "Kinh Duong Vuong",
		Ward:     "12",
		District: "6",
		City:     "Ho Chi Minh",
		Country:  "Viet Nam",
	}

	profile := &Profiles{
		ProfileID:      1,
		ProfileName:    "Personal Profile",
		FirstName:      "Thong",
		LastName:       "Hua Van",
		Email:          "Badger.Smith@gmail.com",
		AccountID:      1,
		Age:            135, // ============> error
		PhoneNumber:    "0908354129",
		DefaultProfile: "Personal Profile",
		FavouriteColor: "#ff0000",
		Addresses:      []*Address{address},
	}

	v := utils.NewValidation()

	err := v.Validate(profile)

	assert.Len(t, err, 1)
}

func TestAbnormalCaseInvalidFormatEmail(t *testing.T) {

	address := &Address{
		Street:   "Kinh Duong Vuong",
		Ward:     "12",
		District: "6",
		City:     "Ho Chi Minh",
		Country:  "Viet Nam",
	}

	profile := &Profiles{
		ProfileID:      1,
		ProfileName:    "Personal Profile",
		FirstName:      "Thong",
		LastName:       "Hua Van",
		Email:          "Badger.Smith@@gmail.com", // ============> error
		AccountID:      1,
		Age:            30,
		PhoneNumber:    "0908354129",
		DefaultProfile: "Personal Profile",
		FavouriteColor: "#ff0000",
		Addresses:      []*Address{address},
	}

	v := utils.NewValidation()

	err := v.Validate(profile)

	assert.Len(t, err, 1)
}

func TestAbnormalCaseInvalidFormatColor(t *testing.T) {

	address := &Address{
		Street:   "Kinh Duong Vuong",
		Ward:     "12",
		District: "6",
		City:     "Ho Chi Minh",
		Country:  "Viet Nam",
	}

	profile := &Profiles{
		ProfileID:      1,
		ProfileName:    "Personal Profile",
		FirstName:      "Thong",
		LastName:       "Hua Van",
		Email:          "Badger.Smith@gmail.com",
		AccountID:      1,
		Age:            30,
		PhoneNumber:    "0908354129",
		DefaultProfile: "Personal Profile",
		FavouriteColor: "#123214213asdsad", // ============> error
		Addresses:      []*Address{address},
	}

	v := utils.NewValidation()

	err := v.Validate(profile)

	assert.Len(t, err, 1)
}

func TestAbnormalCaseMissingCountryReturnsErr(t *testing.T) {

	address := &Address{
		Street:   "Kinh Duong Vuong",
		Ward:     "12",
		District: "6",
		City:     "Ho Chi Minh",
		// Country:  "Viet Nam", // ============> error
	}

	// profile := &Profiles{
	// 	ProfileID:      1,
	// 	ProfileName:    "Personal Profile",
	// 	FirstName:      "Thong",
	// 	LastName:       "Hua Van",
	// 	Email:          "Badger.Smith@gmail.com",
	// 	AccountID:      1,
	// 	Age:            30,
	// 	PhoneNumber:    "0908354129",
	// 	DefaultProfile: "Personal Profile",
	// 	FavouriteColor: "#ff0000",
	// 	Addresses:      []*Address{address},
	// }

	v := utils.NewValidation()

	err := v.Validate(address)

	assert.Len(t, err, 1)
}

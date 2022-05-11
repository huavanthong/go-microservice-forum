/*
 * @File: utils.utils.go
 * @Description: Reusable stuffs for services
 * @Author: Hua Van Thong (seedotech@gmail.com)
 */
package utils

import (
	"errors"

	"github.com/huavanthong/microservice-golang/user-api/common"
	"gopkg.in/mgo.v2/bson"
)

type Utils struct {
}

// ValidateObjectID checks the given ID if it's an object id or not
func (u *Utils) ValidateObjectID(id string) error {
	if bson.IsObjectIdHex(id) != true {
		return errors.New(common.ErrNotObjectIDHex)
	}

	return nil
}

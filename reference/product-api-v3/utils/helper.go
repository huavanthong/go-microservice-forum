package utils

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

// Get type of variable
func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// Get type of struct in models
func TypeOfModel(v interface{}) string {

	// example value: *models.product_phone
	model := fmt.Sprintf("%T", v)

	// match, err := regexp.MatchString("\\*[a-z].[a-z]", model)
	// if err == nil {
	// 	fmt.Println("Match:", match)
	// } else {
	// 	fmt.Println("Error:", err)
	// }
	typeModel := strings.Split(model, ".")

	if len(typeModel) > 2 {
		return "Undefine"
	}

	return typeModel[1]
}

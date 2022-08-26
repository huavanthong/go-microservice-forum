/*
 * @File: category.payload.req.go
 * @Description: Return payload info for product
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

type RequestCreateCategory struct {
	Name        string `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"phone"`
	Description string `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"products relalated to phone category"`
	SubCategory string `json:"subcategory" bson:"subcategory" binding:"required,gt=0,lt=255"`
}

type RequestUpdateCategory struct {
	Name        string `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"phone"`
	SubCategory string `json:"subcategory" bson:"subcategory" binding:"required,gt=0,lt=255"`
	Description string `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"products relalated to phone category"`
}

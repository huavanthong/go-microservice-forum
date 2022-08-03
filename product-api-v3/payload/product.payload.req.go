/*
 * @File: product.payload.req.go
 * @Description: Return payload info for product
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

// Error defines the response error
type RequestCreateProduct struct {
	Name        string `json:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    string `json:"category" binding:"required,gt=0,lt=255" example:"Phone"`
	Summary     string `json:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string `json:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string `json:"imageFile" binding:"required" example:"default.png"`
	Price       string `json:"price" binding:"required,min=0.01" example:"1400$"`
}

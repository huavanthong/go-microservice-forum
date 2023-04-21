package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/response"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Gọi hàm Next() để tiếp tục thực hiện các middleware khác và xử lý request
		c.Next()

		// Kiểm tra xem response có phải là SuccessResponse hay không
		if _, ok := c.Get("success_response"); ok {
			// Nếu là SuccessResponse, lấy dữ liệu và trả về response thành công
			data := c.MustGet("success_response").(response.SuccessResponse)
			c.JSON(data.Status, data)
		} else if _, ok := c.Get("error_response"); ok {
			// Nếu là ErrorResponse, lấy thông tin lỗi và trả về response lỗi
			err := c.MustGet("error_response").(response.ErrorResponse)
			c.JSON(err.Status, err)
		}
	}
}

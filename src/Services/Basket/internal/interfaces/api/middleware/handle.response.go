package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/response"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call Next() for other middleware to handle the request
		c.Next()

		// Check if the response context key exists
		if value, ok := c.Get("response"); ok {
			// If the response is of type *response.Response
			if resp, ok := value.(*response.Response); ok {
				if resp.Success {
					// If the response indicates success, return success status
					c.JSON(http.StatusOK, resp)
				} else {
					// If the response indicates an error, return error status and details
					c.JSON(http.StatusInternalServerError, resp)
				}
				return
			}
		}

		// If no response is found, return a default error message
		c.JSON(http.StatusInternalServerError, response.NewResponse(nil, &response.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Unexpected error occurred",
		}))
	}
}

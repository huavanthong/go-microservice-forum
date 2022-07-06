package controllers

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/huavanthong/microservice-golang/user-api/common"
// 	"github.com/huavanthong/microservice-golang/user-api/models"
// 	"github.com/huavanthong/microservice-golang/user-api/payload"
// 	"github.com/huavanthong/microservice-golang/user-api/utils"
// )

// // MiddlewareValidateProduct validates the product in the request and calls next if ok
// func (u *User) MiddlewareValidateUser() gin.HandlerFunc {

// 	return func(ctx *gin.Context) {

// 		// convert data user request to object user model
// 		var user models.User
// 		if err := ctx.ShouldBindJSON(&user); err != nil {
// 			ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, err.Error()})
// 			return
// 		}

// 		fmt.Println("Check 3")
// 		// validate the product
// 		v := utils.NewValidation()
// 		errs := v.Validate(user)

// 		if len(errs) != 0 {
// 			fmt.Println("Check 4")
// 			// return the validation messages as an array
// 			respondWithError(ctx, 401, "Validating User error")
// 			return
// 		}

// 		ctx.Next()
// 		fmt.Println("Check 5")
// 	}

// }

// func respondWithError(c *gin.Context, code int, message interface{}) {
// 	c.AbortWithStatusJSON(code, gin.H{"error": message})
// }

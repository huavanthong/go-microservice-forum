package controllers

import (
	"fmt"
	"net/http"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"

	"github.com/gin-gonic/gin"
)

type BasketController struct {
	basketService *services.BasketService
}

func NewBasketController(basketService *services.BasketService) *BasketController {
	return &BasketController{
		basketService: basketService,
	}
}

// GetBasket godoc
// @Summary Get basket by user name
// @Description Get basket by user name
// @Tags basket
// @Accept  json
// @Produce  json
// @Param userName path string true "userName"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Success 200 {array} string
// @Router /basket [get]
// GetBasket get basket by user name
func (bc *BasketController) GetBasket(ctx *gin.Context) {

	userName := ctx.Param("userName")
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Name"})
	// 	return
	// }

	basket, err := bc.basketService.GetBasket(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get basket"})
		return
	}

	ctx.JSON(http.StatusOK, basket)
}

// CreateBasket godoc
// @Summary Create basket by user name
// @Description Create basket by user name
// @Tags basket
// @Accept  json
// @Produce  json
// @Param userName path string true "User Name"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Success 200 {array} string
// @Router /basket [post]
// CreateBasket create basket by user name
func (bc *BasketController) CreateBasket(ctx *gin.Context) {

	userName := ctx.Param("userName")
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 	return
	// }

	fmt.Println("Check 1: ", userName)

	shoppingCart, err := bc.basketService.CreateBasket(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create basket"})
		return
	}

	ctx.JSON(http.StatusCreated, shoppingCart)
}

// UpdateBasket godoc
// @Summary Update basket by user name
// @Description Update basket by user name
// @Tags basket
// @Accept  json
// @Produce  json
// @Param userName path string true "userName"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Success 200 {array} string
// @Router /basket [patch]
// UpdateBasket update basket by user name
func (bc *BasketController) UpdateBasket(ctx *gin.Context) {

	userName := ctx.Param("userName")
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Name"})
	// 	return
	// }

	var basket entities.ShoppingCart
	if err := ctx.ShouldBindJSON(&basket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Communicate with Discount.Grpc and calculate latest prices of products into sc
	// for _, item := range basket.Items {
	// 	coupon, err := c.discountGrpc.GetDiscount(item.ProductName)
	// 	if err == nil {
	// 		item.Price -= coupon.Amount
	// 	}
	// }

	if updatedBasket, err := bc.basketService.UpdateBasket(userName, &basket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, updatedBasket)
	}

}

// DeleteBasket godoc
// @Summary Delete basket by user name
// @Description Delete basket by user name
// @Tags basket
// @Accept  json
// @Produce  json
// @Param userName path string true "userName"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Success 200 {array} string
// @Router /basket/{userName} [delete]
// UpdateBasket update basket by user name
func (bc *BasketController) DeleteBasket(ctx *gin.Context) {

	userName := ctx.Param("userName")
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Name"})
	// 	return
	// }

	if err := bc.basketService.DeleteBasket(userName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

// Checkout godoc
// @Summary Checkout basket after completed items in shopping cart
// @Description Checkout basket
// @Tags basket
// @Accept  json
// @Produce  json
// @Param basketCheckout body string true "Basket Checkout"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Success 200 {array} string
// @Router /Basket/{userName} [get]
// Checkout checkout basket
func (bc *BasketController) Checkout(ctx *gin.Context) {
	var basketCheckout entities.BasketCheckout
	if err := ctx.ShouldBindJSON(&basketCheckout); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get existing basket with total price
	basket, err := bc.basketService.GetBasket(basketCheckout.UserName)
	if err != nil || basket == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "basket not found"})
		return
	}

	// remove the basket
	if err := bc.basketService.DeleteBasket(basket.UserName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

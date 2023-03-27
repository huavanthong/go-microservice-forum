package controllers

import (
	"net/http"
	"shopping-cart-service/internal/domain/entities"
	"shopping-cart-service/internal/domain/repositories"

	"github.com/gin-gonic/gin"
)

type BasketController struct {
	repo repositories.BasketRepository
	/* Mở rộng sau:
	discountGrpc    DiscountGrpcService
	*/
}

func NewBasketController(repo repositories.BasketRepository) *BasketController {
	return &BasketController{repo: repo}
}

func (c *BasketController) GetBasket(ctx *gin.Context) {
	userName := ctx.Param("userName")
	basket, err := c.repo.GetBasket(userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if basket == nil {
		basket = entities.NewShoppingCart(userName)
	}
	ctx.JSON(http.StatusOK, basket)
}

func (c *BasketController) UpdateBasket(ctx *gin.Context) {
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

	if updatedBasket, err := c.repo.UpdateBasket(&basket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, updatedBasket)
	}

}

func (c *BasketController) DeleteBasket(ctx *gin.Context) {
	userName := ctx.Param("userName")
	if err := c.repo.DeleteBasket(userName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func (c *BasketController) Checkout(ctx *gin.Context) {
	var basketCheckout entities.BasketCheckout
	if err := ctx.ShouldBindJSON(&basketCheckout); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get existing basket with total price
	basket, err := c.repo.GetBasket(basketCheckout.UserName)
	if err != nil || basket == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "basket not found"})
		return
	}

	// remove the basket
	if err := c.repo.DeleteBasket(basket.UserName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

package controllers

import (
	"net/http"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/response"

	"github.com/gin-gonic/gin"
)

type BasketController struct {
	basketService services.BasketService
}

func NewBasketController(basketService services.BasketService) BasketController {
	return BasketController{
		basketService: basketService,
	}
}

// GetBasket godoc
// @Summary Get basket by user id
// @Description Get basket by user id
// @Tags basket
// @Accept  json
// @Produce  json
// @Param userid path string true "User ID"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Success 200 {array} response.SuccessResponse
// @Router /basket/{userid} [get]
func (bc *BasketController) GetBasket(ctx *gin.Context) {

	userId := ctx.Param("userid")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "User ID is required"))
		return
	}

	basket, err := bc.basketService.GetBasket(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, "Basket not exists"))
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(basket))
	return
}

// CreateBasket godoc
// @Summary Create basket by user id
// @Description Create basket by user id
// @Tags basket
// @Accept  json
// @Produce  json
// @Param basket body models.CreateBasketRequest true "New Basket"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Success 200 {object} response.SuccessResponse
// @Router /basket [post]
// CreateBasket create basket by user id
func (bc *BasketController) CreateBasket(ctx *gin.Context) {

	// Deserialization data from request
	var cbq models.CreateBasketRequest
	if err := ctx.ShouldBindJSON(&cbq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	basket, err := bc.basketService.CreateBasket(&cbq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, "Failed to create basket"))
		return
	}

	ctx.JSON(http.StatusCreated, response.NewSuccessResponse(basket))
	return
}

// UpdateBasket godoc
// @Summary Update basket by user id
// @Description Update basket by user id
// @Tags basket
// @Accept  json
// @Produce  json
// @Param basket body models.UpdateBasketRequest true "Update Basket"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Success 200 {object} response.SuccessResponse
// @Router /basket [patch]
func (bc *BasketController) UpdateBasket(ctx *gin.Context) {

	// Deserialization data from request
	var ubq models.UpdateBasketRequest
	if err := ctx.ShouldBindJSON(&ubq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// Communicate with Discount.Grpc and calculate latest prices of products into sc
	// for _, item := range basket.Items {
	// 	coupon, err := c.discountGrpc.GetDiscount(item.ProductName)
	// 	if err == nil {
	// 		item.Price -= coupon.Amount
	// 	}
	// }
	updatedBasket, err := bc.basketService.UpdateBasket(&ubq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(updatedBasket))
	return
}

// DeleteBasket godoc
// @Summary Delete basket by user id
// @Description Delete basket by user id
// @Tags basket
// @Accept  json
// @Produce  json
// @Param userid path string true "User ID"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Success 200 {object} response.SuccessResponse
// @Router /basket/{userid} [delete]
func (bc *BasketController) DeleteBasket(ctx *gin.Context) {

	userId := ctx.Param("userid")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "User ID is required"))
		return
	}

	if err := bc.basketService.DeleteBasket(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse("Deleted successfully"))

}

// Checkout godoc
// @Summary Checkout basket after completed items in shopping cart
// @Description Checkout basket
// @Tags basket
// @Accept  json
// @Produce  json
// @Param basketCheckout body string true "Basket Checkout"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Success 200 {object} response.SuccessResponse
// @Router /basket/checkout [post]
// Checkout checkout basket
func (bc *BasketController) Checkout(ctx *gin.Context) {

	// Deserialization data from request
	var basketCheckout entities.BasketCheckout
	if err := ctx.ShouldBindJSON(&basketCheckout); err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// get existing basket with total price
	basket, err := bc.basketService.GetBasket(basketCheckout.UserID)
	if err != nil || basket == nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "basket not found"))
		return
	}

	// remove the basket
	if err := bc.basketService.DeleteBasket(basket.UserName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/services"
)

type DiscountController struct {
	discountService services.DiscountService
}

func NewDiscountController(discountService services.DiscountService) *DiscountController {
	return &DiscountController{
		discountService: discountService,
	}
}

// GetDiscount godoc
// @Summary Get discount for product name
// @Description Get discount for product name
// @Tags discount
// @Accept  json
// @Produce  json
// @Success 200 {object} payload.UserRegisterSuccess
// @Failure 400 {object} models.GenericResponse
// @Router /discount/:productName [get]
func (c *DiscountController) GetDiscount(ctx *gin.Context) {

	var reqDiscount *models.GetDiscountRequest

	// from context, bind user info to json
	if err := ctx.ShouldBindJSON(&reqDiscount); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "Invalid data request",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	discount, err := c.discountService.GetDiscountByID(reqDiscount.ID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			ctx.JSON(http.StatusNotFound,
				models.GenericResponse{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}

		ctx.JSON(http.StatusOK,
			models.GenericResponse{
				Success: true,
				Code:    http.StatusOK,
				Message: "Get discount success",
				Data:    models.GetDiscountResponse,
				Errors:  nil,
			})
		return
	}

	ctx.JSON(http.StatusOK, discount)
	return
}

// CreateDiscount godoc
// @Summary Create coupon
// @Description Create coupon
// @Tags discount
// @Accept  json
// @Produce  json
// @Success 200 {object} http.StatusOK
// @Failure 400 {object} http.StatusBadRequest
// @Success 500 {object} http.StatusInternalServerError
// @Router /discount [post]
func (c *DiscountController) CreateDiscount(ctx *gin.Context) {
	var coupon models.Coupon
	if err := ctx.ShouldBindJSON(&coupon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := c.discountService.CreateDiscount(coupon)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create discount"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "discount created successfully"})
}

// UpdateDiscount godoc
// @Summary Update coupon
// @Description Update coupon
// @Tags discount
// @Accept  json
// @Produce  json
// @Success 200 {object} http.StatusOK
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
// @Success 500 {object} http.StatusInternalServerError
// @Router /discount [put]
func (c *DiscountController) UpdateDiscount(ctx *gin.Context) {
	var coupon models.Coupon
	if err := ctx.ShouldBindJSON(&coupon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := c.discountService.UpdateDiscount(coupon)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "discount updated successfully"})
}

// DeleteDiscount godoc
// @Summary Delete coupon by product name
// @Description Delete coupon by product name
// @Tags discount
// @Accept  json
// @Produce  json
// @Success 200 {object} http.StatusOK
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
// @Success 500 {object} http.StatusInternalServerError
// @Router /discount [put]
func (c *DiscountController) DeleteDiscount(ctx *gin.Context) {
	productName := ctx.Query("productName")
	if productName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing productName parameter"})
		return
	}

	ok, err := c.discountService.DeleteDiscount(productName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "discount deleted successfully"})
}

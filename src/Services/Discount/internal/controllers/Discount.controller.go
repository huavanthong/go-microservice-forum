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
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
// @Success 500 {object} http.StatusInternalServerError
// @Router /discount/:productName [get]
func (c *DiscountController) GetDiscount(cxt *gin.Context) {
	productName := cxt.Query("productName")
	if productName == "" {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "missing productName parameter"})
		return
	}

	discount, err := c.discountService.GetDiscount(productName)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			cxt.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, discount)
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
func (c *DiscountController) CreateDiscount(cxt *gin.Context) {
	var coupon models.Coupon
	if err := cxt.ShouldBindJSON(&coupon); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := c.discountService.CreateDiscount(coupon)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create discount"})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "discount created successfully"})
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
func (c *DiscountController) UpdateDiscount(cxt *gin.Context) {
	var coupon models.Coupon
	if err := cxt.ShouldBindJSON(&coupon); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := c.discountService.UpdateDiscount(coupon)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		cxt.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "discount updated successfully"})
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
func (c *DiscountController) DeleteDiscount(cxt *gin.Context) {
	productName := cxt.Query("productName")
	if productName == "" {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "missing productName parameter"})
		return
	}

	ok, err := c.discountService.DeleteDiscount(productName)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		cxt.JSON(http.StatusNotFound, gin.H{"error": "discount not found"})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "discount deleted successfully"})
}

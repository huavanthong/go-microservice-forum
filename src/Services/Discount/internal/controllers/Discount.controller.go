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

func (c *DiscountController) GetDiscount(cxt *gin.Context) {
	productName := cxt.Query("productName")
	if productName == "" {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "missing productName parameter"})
		return
	}

	coupon, err := c.discountService.GetDiscount(productName)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			cxt.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, coupon)
}

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

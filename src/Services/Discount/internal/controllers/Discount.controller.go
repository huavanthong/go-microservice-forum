package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/services"
)

type DiscountController struct {
	discountService services.DiscountService
}

func NewDiscountController(discountService services.DiscountService) DiscountController {
	return DiscountController{
		discountService: discountService,
	}
}

// GetDiscount godoc
// @Summary Get discount for product name
// @Description Get discount for product name
// @Tags discount
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Failure 404 {object} models.GenericResponse
// @Router /discount/{discountId} [get]
func (c *DiscountController) GetDiscount(ctx *gin.Context) {

	// get user ID from URL path
	discountId := ctx.Param("discountId")

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

	discount, err := c.discountService.GetDiscount(discountId)
	if err != nil {
		// Not found
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				models.GenericResponse{
					Success: false,
					Code:    http.StatusNotFound,
					Message: "Discount not found",
					Data:    nil,
					Errors:  []string{err.Error()},
				})
			return
		}
		// Success
		ctx.JSON(http.StatusOK,
			models.GenericResponse{
				Success: true,
				Code:    http.StatusOK,
				Message: "Get discount success",
				Data:    discount,
				Errors:  nil,
			})
		return
	}

	ctx.JSON(http.StatusOK, discount)
	return
}

// CreateDiscount godoc
// @Summary Create discount for product
// @Description Create discount for product
// @Tags discount
// @Accept  json
// @Produce  json
// @Success 200 {object} http.StatusOK
// @Failure 400 {object} http.StatusBadRequest
// @Success 500 {object} http.StatusInternalServerError
// @Router /discount [post]
func (c *DiscountController) CreateDiscount(ctx *gin.Context) {

	var discount models.Discount
	if err := ctx.ShouldBindJSON(&discount); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "Invalid data request to create discount",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	err := c.discountService.CreateDiscount(&discount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "failed to create discount",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GenericResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "discount created successfully",
			Data:    nil,
			Errors:  nil,
		})
	return
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
	var discount models.Discount
	if err := ctx.ShouldBindJSON(&discount); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "Invalid data request to update discount",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	err := c.discountService.UpdateDiscount(&discount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "Internal server service error",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GenericResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "discount updated successfully",
			Data:    nil,
			Errors:  nil,
		})
	return
}

// DeleteDiscount godoc
// @Summary Delete coupon by product name
// @Description Delete coupon by product name
// @Tags discount
// @Accept  json
// @Produce  json
// @Param discountId path string true "Discout ID"
// @Success 200 {object} http.StatusOK
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
// @Success 500 {object} http.StatusInternalServerError
// @Router /discount/{discountId}} [delete]
func (c *DiscountController) DeleteDiscount(ctx *gin.Context) {

	// get user ID from URL path
	discountId := ctx.Param("discountId")

	// call discount service to find discount by ID
	err := c.discountService.DeleteDiscount(discountId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "Internal server service error",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GenericResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "discount deleted successfully",
			Data:    nil,
			Errors:  nil,
		})
	return
}

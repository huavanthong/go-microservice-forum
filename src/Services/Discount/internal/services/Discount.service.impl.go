package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/repositories"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/utils"
)

// DiscountService represents the discount service
type DiscountServiceImpl struct {
	discountRepo repositories.DiscountRepository
}

func NewDiscountServiceImpl(discountRepo repositories.DiscountRepository) DiscountService {
	return &DiscountServiceImpl{
		discountRepo: discountRepo,
	}
}

// Generate response
func generateGetDiscountResponse(discount *models.Discount) (*models.GetDiscountResponse, error) {

	return &models.GetDiscountResponse{
		ID:           discount.ID,
		ProductName:  discount.ProductName,
		Description:  discount.Description,
		DiscountType: discount.DiscountType,
		Percentage:   discount.Percentage,
		Quantity:     discount.Quantity,
		StartDate:    discount.StartDate,
		EndDate:      discount.EndDate,
		Available:    true,
	}, nil
}

// GetDiscount gets the discount based on the input parameters
func (ds *DiscountServiceImpl) GetDiscount(ID string) (*models.GetDiscountResponse, error) {

	// convert string id to int id
	intId, _ := strconv.Atoi(ID)

	// Get discount from repository
	discount, err := ds.discountRepo.GetDiscount(intId)
	if err != nil {

		return nil, fmt.Errorf("Error while getting discount: %v", err)
	}

	if discount == nil {
		return nil, fmt.Errorf("Discount not found: %v", err)
	}

	return generateGetDiscountResponse(discount)
}

// Convert request create discount to discount entity
func convertRequestCreatetToDiscount(discountReq *models.CreateDiscountRequest) (*models.Discount, error) {

	// Business case specific that only one value for amount or percent
	if discountReq.Amount > 0 && discountReq.Percentage > 0 {
		return nil, fmt.Errorf("Cannot create discount with both amount and discount percent")
	}

	if discountReq.Amount == 0 && discountReq.Percentage == 0 {
		return nil, fmt.Errorf("Cannot create discount with both zero amount and zero discount percent")
	}

	err := utils.ValidateStartDateEndDate(discountReq.StartDate, discountReq.EndDate)
	if err != nil {
		return nil, err
	}

	discount := &models.Discount{
		ProductID:    discountReq.ProductID,
		ProductName:  discountReq.ProductName,
		Description:  discountReq.Description,
		DiscountType: discountReq.DiscountType,
		Percentage:   discountReq.Percentage,
		Amount:       discountReq.Amount,
		Quantity:     discountReq.Quantity,
		StartDate:    discountReq.StartDate,
		EndDate:      discountReq.EndDate,
	}

	return discount, nil
}

func (ds *DiscountServiceImpl) CreateDiscount(discountReq *models.CreateDiscountRequest) (*models.Discount, error) {

	discount, err := convertRequestCreatetToDiscount(discountReq)
	if err != nil {
		return nil, fmt.Errorf("Invalid request to create discount", err)
	}

	return ds.discountRepo.CreateDiscount(discount)
}

// Convert request update discount to discount entity
func convertRequestUpdateToDiscount(discountReq *models.UpdateDiscountRequest) (*models.Discount, error) {

	// Business case specific that only one value for amount or percent
	if discountReq.Amount > 0 && discountReq.Percentage > 0 {
		return nil, fmt.Errorf("Cannot create discount with both amount and discount percent")
	}

	if discountReq.Amount == 0 && discountReq.Percentage == 0 {
		return nil, fmt.Errorf("Cannot create discount with both zero amount and zero discount percent")
	}

	err := utils.ValidateStartDateEndDate(discountReq.StartDate, discountReq.EndDate)
	if err != nil {
		return nil, err
	}

	updateAt := time.Now()

	discount := &models.Discount{
		ProductID:    discountReq.ProductID,
		ProductName:  discountReq.ProductName,
		Description:  discountReq.Description,
		DiscountType: discountReq.DiscountType,
		Percentage:   discountReq.Percentage,
		Amount:       discountReq.Amount,
		Quantity:     discountReq.Quantity,
		StartDate:    discountReq.StartDate,
		EndDate:      discountReq.EndDate,
		UpdatedAt:    updateAt,
	}

	return discount, nil
}

func (ds *DiscountServiceImpl) UpdateDiscount(discountReq *models.UpdateDiscountRequest) (*models.Discount, error) {
	// Convert request update to discount
	discount, err := convertRequestUpdateToDiscount(discountReq)

	if err != nil {
		return nil, fmt.Errorf("Invalid request to update discount", err)
	}

	return ds.discountRepo.UpdateDiscount(discount)
}

func (ds *DiscountServiceImpl) DeleteDiscount(ID string) error {

	// convert string id to int id
	intId, _ := strconv.Atoi(ID)

	return ds.discountRepo.DeleteDiscount(intId)
}

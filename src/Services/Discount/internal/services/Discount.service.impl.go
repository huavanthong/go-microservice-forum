package services

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/repositories"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DiscountService represents the discount service
type DiscountServiceImpl struct {
	log          *zap.Logger
	discountRepo repositories.DiscountRepository
}

func NewDiscountServiceImpl(log *zap.Logger, discountRepo repositories.DiscountRepository) DiscountService {
	return &DiscountServiceImpl{
		log:          log,
		discountRepo: discountRepo,
	}
}

// GetDiscount gets the discount based on the input parameters
func (ds *DiscountServiceImpl) GetDiscount(ID string) (*models.GetDiscountResponse, error) {

	// convert string id to int id
	intId, _ := strconv.Atoi(ID)

	// Get discount from repository
	discountResponse, err := ds.discountRepo.GetDiscount(intId)
	if err != nil {
		fmt.Printf("Error while getting discount: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	if discountResponse == nil {
		return nil, status.Error(codes.NotFound, "Discount not found")
	}

	// Convert to get response
	res := models.FilteredGetResponse(discountResponse, true)

	return res, nil
}

func (ds *DiscountServiceImpl) CreateDiscount(discount *models.Discount) error {

	return ds.discountRepo.CreateDiscount(discount)
}

func (ds *DiscountServiceImpl) UpdateDiscount(discount *models.Discount) error {

	return ds.discountRepo.UpdateDiscount(discount)
}

func (ds *DiscountServiceImpl) DeleteDiscount(ID string) error {

	// convert string id to int id
	intId, _ := strconv.Atoi(ID)

	return ds.discountRepo.DeleteDiscount(intId)
}

package services

import (
	"fmt"
	"log"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/repositories"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DiscountService represents the discount service
type DiscountService struct {
	logger       *log.Logger
	discountRepo repositories.DiscountRepository
}

func NewDiscountService(logger *log.Logger, discountRepo repositories.DiscountRepository) *DiscountService {
	return &DiscountService{
		logger:       logger,
		discountRepo: discountRepo,
	}
}

// GetDiscount gets the discount based on the input parameters
func (s *DiscountService) GetDiscountByID(ID int) (*models.GetDiscountResponse, error) {

	// Get discount from repository
	discountResponse, err := s.discountRepo.GetDiscountByID(ID)
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

func (s *DiscountService) CreateDiscount(discount *models.Discount) error {

	return s.discountRepo.CreateDiscount(discount)
}

func (s *DiscountService) UpdateDiscount(discount *models.Discount) error {

	return s.discountRepo.UpdateDiscount(discount)
}

func (s *DiscountService) DeleteDiscount(ID int) error {

	return s.discountRepo.DeleteDiscountByID(ID)
}

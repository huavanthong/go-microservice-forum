package services

import (
	"context"
	"errors"
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
func (s *DiscountService) GetDiscountByID(ctx context.Context, req *models.GetDiscountRequest) (*models.GetDiscountResponse, error) {

	// Get discount from repository
	discount, err := s.discountRepo.GetDiscount(req.ProductName)
	if err != nil {
		fmt.Printf("Error while getting discount: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	if discount == nil {
		return nil, status.Error(codes.NotFound, "Coupon not found")
	}

	discount := &models.Discount{
		ProductName: coupon.ProductName,
		Amount:      coupon.Amount,
		Description: coupon.Description,
	}

	res := &models.DiscountResponse{
		Discount: discount,
	}

	return res, nil
}
func (s *DiscountService) CreateDiscount(ctx context.Context, req *models.DiscountRequest) (*models.DiscountResponse, error) {

	coupon := &models.Coupon{
		ProductName: req.ProductName,
		Description: req.Description,
		Amount:      req.Amount,
	}

	if err := s.discountRepo.CreateDiscount(coupon); err != nil {
		fmt.Printf("Error while creating discount: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	discount := &models.Discount{
		ProductName: coupon.ProductName,
		Amount:      coupon.Amount,
		Description: coupon.Description,
	}

	res := &models.DiscountResponse{
		Discount: discount,
	}

	return res, nil
}

func (s *DiscountService) UpdateDiscount(ctx context.Context, req *models.Discount) (*models.DiscountResponse, error) {
	coupon := &models.Coupon{
		Id:          req.Id,
		ProductName: req.ProductName,
		Description: req.Description,
		Amount:      req.Amount,
	}

	if err := s.discountRepo.UpdateDiscount(coupon); err != nil {
		fmt.Printf("Error while updating discount: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	discount := &models.Discount{
		ProductName: coupon.ProductName,
		Amount:      coupon.Amount,
		Description: coupon.Description,
	}

	res := &models.DiscountResponse{
		Discount: discount,
	}

	return res, nil
}

func (s *DiscountService) DeleteDiscount(ctx context.Context, req *models.DiscountRequest) (*models.DeleteDiscountResponse, error) {

	if err := s.discountRepo.DeleteDiscount(req.ProductName); err != nil {
		fmt.Printf("Error while deleting discount: %v\n", err)
		if errors.Is(err, repositories.ErrCouponNotFound) {
			return nil, status.Error(codes.NotFound, "Coupon not found")
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res := &models.DeleteDiscountResponse{
		Success: true,
	}

	return res, nil
}

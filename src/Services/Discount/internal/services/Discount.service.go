package discount

import (
	"context"
	"errors"
	"fmt"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/config"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/model"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DiscountService struct {
	repo   repository.DiscountRepository
	config *config.Configuration
}

func NewDiscountService(repo repository.DiscountRepository, config *config.Configuration) *DiscountService {
	return &DiscountService{
		repo:   repo,
		config: config,
	}
}

func (s *DiscountService) GetDiscount(ctx context.Context, req *model.DiscountRequest) (*model.DiscountResponse, error) {
	coupon, err := s.repo.GetDiscount(req.ProductName)
	if err != nil {
		fmt.Printf("Error while getting discount: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	if coupon == nil {
		return nil, status.Error(codes.NotFound, "Coupon not found")
	}

	discount := &model.Discount{
		ProductName: coupon.ProductName,
		Amount:      coupon.Amount,
		Description: coupon.Description,
	}

	res := &model.DiscountResponse{
		Discount: discount,
	}

	return res, nil
}

func (s *DiscountService) CreateDiscount(ctx context.Context, req *model.Discount) (*model.DiscountResponse, error) {
	coupon := &model.Coupon{
		ProductName: req.ProductName,
		Description: req.Description,
		Amount:      req.Amount,
	}

	if err := s.repo.CreateDiscount(coupon); err != nil {
		fmt.Printf("Error while creating discount: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	discount := &model.Discount{
		ProductName: coupon.ProductName,
		Amount:      coupon.Amount,
		Description: coupon.Description,
	}

	res := &model.DiscountResponse{
		Discount: discount,
	}

	return res, nil
}

func (s *DiscountService) UpdateDiscount(ctx context.Context, req *model.Discount) (*model.DiscountResponse, error) {
	coupon := &model.Coupon{
		Id:          req.Id,
		ProductName: req.ProductName,
		Description: req.Description,
		Amount:      req.Amount,
	}

	if err := s.repo.UpdateDiscount(coupon); err != nil {
		fmt.Printf("Error while updating discount: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	discount := &model.Discount{
		ProductName: coupon.ProductName,
		Amount:      coupon.Amount,
		Description: coupon.Description,
	}

	res := &model.DiscountResponse{
		Discount: discount,
	}

	return res, nil
}

func (s *DiscountService) DeleteDiscount(ctx context.Context, req *model.DiscountRequest) (*model.DeleteDiscountResponse, error) {
	if err := s.repo.DeleteDiscount(req.ProductName); err != nil {
		fmt.Printf("Error while deleting discount: %v\n", err)
		if errors.Is(err, repository.ErrCouponNotFound) {
			return nil, status.Error(codes.NotFound, "Coupon not found")
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res := &model.DeleteDiscountResponse{
		Success: true,
	}

	return res, nil
}

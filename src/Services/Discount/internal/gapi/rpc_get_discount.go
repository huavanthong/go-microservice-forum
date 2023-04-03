package gapi

import (
	"context"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/proto/discountpb"
)

// GetDiscount gets the discount based on the input parameters
func (server *Server) GetDiscount(ctx context.Context, req *discountpb.SignUpUserInput) (*discountpb.GenericResponse, error) {

	// Validate request
	if req.GetProductId() == "" {
		return nil, status.Error(codes.InvalidArgument, "product ID cannot be empty")
	}
	if req.GetQuantity() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "quantity must be greater than 0")
	}

	// Get discount from repository
	discount, err := s.discountRepo.GetDiscount(ctx, req.GetProductId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get discount")
	}

	// Check if discount is available
	if discount == nil {
		return &discountpb.DiscountResponse{
			Available: false,
		}, nil
	}

	// Calculate total discount
	totalDiscount := float32(0)
	for _, d := range discount.Tiers {
		if req.GetQuantity() >= d.Quantity {
			totalDiscount += d.Discount
		}
	}

	// Build the response
	res := &discountpb.DiscountResponse{
		DiscountId: discount.DiscountID,
		Amount:     discount.Amount,
	}

	return res, nil
}

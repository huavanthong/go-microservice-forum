package services

// func (s *DiscountService) GetCoupon(ctx context.Context, req *discountpb.CouponRequest) (*discountpb.CouponResponse, error) {
// 	// Validate request
// 	if req.GetCouponCode() == "" {
// 		return nil, status.Error(codes.InvalidArgument, "coupon code cannot be empty")
// 	}

// 	// Get coupon from repository
// 	coupon, err := s.couponRepo.GetCoupon(ctx, req.GetCouponCode())
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "failed to get coupon")
// 	}

// 	// Check if coupon is available
// 	if coupon == nil {
// 		return &discountpb.CouponResponse{
// 			Available: false,
// 		}, nil
// 	}

// 	// Return coupon response
// 	return &discountpb.CouponResponse{
// 		Available: true,
// 		Discount:  coupon.Discount,
// 	}, nil
// }

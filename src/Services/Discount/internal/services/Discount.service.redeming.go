package services

/*
func (s *BasketService) RedeemingCoupon(ctx context.Context, req *pb.RedeemCouponRequest) (*pb.RedeemCouponResponse, error) {
	// Get the basket
	basket, err := s.basketRepo.GetBasket(ctx, req.BasketId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get basket: %v", err)
	} // Check if the basket has been checked out
	if basket.CheckedOut {
		return nil, status.Error(codes.InvalidArgument, "basket has already been checked out")
	}

	// Get the coupon
	coupon, err := s.couponRepo.GetCouponByCode(ctx, req.CouponCode)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get coupon: %v", err)
	}

	// Check if the coupon has expired
	if coupon.ExpirationDate.Before(time.Now()) {
		return nil, status.Error(codes.InvalidArgument, "coupon has expired")
	}

	// Check if the coupon has been used
	if coupon.Used {
		return nil, status.Error(codes.InvalidArgument, "coupon has already been used")
	}

	// Check if the coupon can be used with the current basket
	if coupon.MinBasketAmount > basket.TotalPrice {
		return nil, status.Error(codes.InvalidArgument, "coupon cannot be used with this basket")
	}

	// Apply the coupon to the basket
	basket.Discount = &pb.Discount{
		Code:        coupon.Code,
		Description: coupon.Description,
		Amount:      coupon.Amount,
	}

	// Mark the coupon as used
	coupon.Used = true
	err = s.couponRepo.UpdateCoupon(ctx, coupon)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update coupon: %v", err)
	}

	// Update the basket
	err = s.basketRepo.UpdateBasket(ctx, basket)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update basket: %v", err)
	}

	return &pb.RedeemCouponResponse{
		Success: true,
	}, nil

}
*/

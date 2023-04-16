package handler

import (
	"context"
	"errors"
)

type RequestOTPCommand struct {
	CustomerID  string
	PhoneNumber string
}

type RequestOTPCommandHandler struct {
	// dependencies
	otpService OtpService
	eventBus   EventBus
}

func (h *RequestOTPCommandHandler) Handle(ctx context.Context, cmd RequestOTPCommand) error {
	// validate input
	if cmd.CustomerID == "" || cmd.PhoneNumber == "" {
		return errors.New("invalid input")
	}

	// request OTP from OTP service
	otp, err := h.otpService.RequestOTP(ctx, cmd.PhoneNumber)
	if err != nil {
		return err
	}

	// send OTP to customer via SMS or email
	// ...

	// publish event to confirm OTP request
	event := &OTPRequestedEvent{
		CustomerID: cmd.CustomerID,
	}
	return h.eventBus.Publish(ctx, event)
}

package utils

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrInvalidOTP       = errors.New("invalid OTP")
	ErrFailedToSendOTP  = errors.New("failed to send OTP")
	ErrFailedToStoreOTP = errors.New("failed to store OTP")
	ErrFailedToFetchOTP = errors.New("failed to fetch OTP")
	ErrOTPExpired       = errors.New("OTP expired")
)

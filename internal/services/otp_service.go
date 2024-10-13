package services

import (
	"github.com/sirupsen/logrus"
	"otp-auth/internal/interfaces"
	"otp-auth/internal/repositories"
	"otp-auth/internal/utils"
	"time"
)

type OtpService struct {
	cache       *repositories.CacheRepository
	otpProvider interfaces.OtpProvider
}

func NewOtpService(cacheRepo *repositories.CacheRepository, otpProvider interfaces.OtpProvider) *OtpService {
	return &OtpService{
		cache:       cacheRepo,
		otpProvider: otpProvider,
	}
}

func (s *OtpService) GenerateAndSendOtp(phoneNumber string) error {

	otp := s.otpProvider.GenerateOtp()
	utils.LogInfo("Generated OTP", logrus.Fields{
		"phone_number": phoneNumber,
		"otp":          otp,
	})

	err := s.cache.SaveOtp(phoneNumber, otp, 5*time.Minute)
	if err != nil {
		utils.LogError("Failed to store OTP in Redis", err, logrus.Fields{
			"phone_number": phoneNumber,
		})
		return utils.ErrFailedToStoreOTP
	}

	err = s.otpProvider.SendOtp(phoneNumber, otp)
	if err != nil {
		utils.LogError("Failed to send OTP in Redis", err, logrus.Fields{
			"phone_number": phoneNumber,
		})
		return err
	}

	utils.LogInfo("OTP successfully generated and sent", logrus.Fields{
		"phone_number": phoneNumber,
		"otp":          otp,
	})

	return nil
}

func (s *OtpService) VerifyOtp(phoneNumber string, otp string) error {

	cachedOtp, err := s.cache.GetOtp(phoneNumber)
	if err != nil {
		utils.LogError("Failed to fetch OTP from Redis", err, logrus.Fields{
			"phone_number": phoneNumber,
		})
	}

	if cachedOtp != otp {
		utils.LogError("Failed to verify OTP from Redis", err, logrus.Fields{
			"phone_number": phoneNumber,
			"otp":          otp,
		})
		return utils.ErrInvalidOTP
	}

	utils.LogInfo("OTP verified successfully", logrus.Fields{
		"phone_number": phoneNumber,
	})

	return nil
}

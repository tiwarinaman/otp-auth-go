package repositories

import (
	"otp-auth/internal/utils"
	"time"
)

type CacheRepository struct {
}

func NewCacheRepository() *CacheRepository {
	return &CacheRepository{}
}

func (c *CacheRepository) SaveOtp(phoneNumber string, otp string, ttl time.Duration) error {

	err := utils.SetValue(phoneNumber, otp, ttl)
	if err != nil {
		utils.LogError("Failed to store OTP in Redis", err, map[string]interface{}{
			"phone_number": phoneNumber,
		})
		return utils.ErrFailedToStoreOTP
	}

	utils.LogInfo("OTP stored in Redis", map[string]interface{}{
		"phone_number": phoneNumber,
		"otp":          otp,
	})

	return nil
}

func (c *CacheRepository) GetOtp(phoneNumber string) (string, error) {

	otp, err := utils.GetValue(phoneNumber)
	if err != nil {
		utils.LogError("Failed to get the otp", err, map[string]interface{}{
			"phone_number": phoneNumber,
		})
		return "", err
	}

	return otp, nil
}

package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"otp-auth/internal/utils"
	"time"
)

type SMSOTPProvider struct {
	rng *rand.Rand
}

func NewSmsOtpProvider() *SMSOTPProvider {
	source := rand.NewSource(time.Now().UnixNano())
	return &SMSOTPProvider{
		rng: rand.New(source),
	}
}

func (p *SMSOTPProvider) GenerateOtp() string {
	return fmt.Sprintf("%06d", p.rng.Intn(1000000))
}

func (p *SMSOTPProvider) SendOtp(phoneNumber string, otp string) error {
	err := utils.SendSMS(phoneNumber, otp)
	if err != nil {
		utils.LogError("Failed to send OTP via SMS", err, logrus.Fields{
			"phone_number": phoneNumber,
		})
		return utils.ErrFailedToSendOTP
	}
	utils.LogInfo("OTP sent successfully via SMS", logrus.Fields{
		"phone_number": phoneNumber,
		"otp":          otp,
	})
	return nil
}

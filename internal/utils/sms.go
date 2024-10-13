package utils

import (
	"github.com/sirupsen/logrus"
)

func SendSMS(phoneNumber string, otp string) error {
	LogInfo("Sending otp", logrus.Fields{
		"phone_number": phoneNumber,
	})
	return nil
}

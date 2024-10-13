package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"otp-auth/internal/constants"
	"otp-auth/internal/services"
	"otp-auth/internal/utils"
)

type AuthController struct {
	otpService *services.OtpService
}

func NewAuthController(otpService *services.OtpService) *AuthController {
	return &AuthController{
		otpService: otpService,
	}
}

func (ctrl *AuthController) RequestOtp(ctx *gin.Context) {

	requestId := ctx.GetHeader(constants.XRequestId)

	var req struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
	}

	utils.LogInfo("Received OTP request", logrus.Fields{
		"request_id":   requestId,
		"phone_number": req.PhoneNumber,
		"operation":    "RequestOTP",
	})

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.LogError("Invalid request data", err, logrus.Fields{
			"request_id": requestId,
			"operation":  "RequestOTP",
		})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.otpService.GenerateAndSendOtp(req.PhoneNumber); err != nil {
		utils.LogError("Failed to generate or send OTP", err, logrus.Fields{
			"request_id":   requestId,
			"phone_number": req.PhoneNumber,
			"operation":    "RequestOTP",
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.LogInfo("OTP successfully sent", logrus.Fields{
		"request_id":   requestId,
		"phone_number": req.PhoneNumber,
		"operation":    "RequestOTP",
	})

	ctx.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func (ctrl *AuthController) VerifyOTP(ctx *gin.Context) {

	requestId := ctx.GetString(constants.XRequestId)

	var req struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		OTP         string `json:"otp" binding:"required"`
	}

	utils.LogInfo("Received OTP verification request", logrus.Fields{
		"request_id":   requestId,
		"phone_number": req.PhoneNumber,
		"operation":    "VerifyOTP",
	})

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.LogError("Invalid verification request data", err, logrus.Fields{
			"request_id": requestId,
			"operation":  "VerifyOTP",
		})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.otpService.VerifyOtp(req.PhoneNumber, req.OTP); err != nil {
		utils.LogError("Failed to verify OTP", err, logrus.Fields{
			"request_id":   requestId,
			"phone_number": req.PhoneNumber,
			"operation":    "VerifyOTP",
		})
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	utils.LogInfo("OTP verified successfully", logrus.Fields{
		"request_id":   requestId,
		"phone_number": req.PhoneNumber,
		"operation":    "VerifyOTP",
	})

	ctx.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}

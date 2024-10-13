package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"otp-auth/configs"
	"otp-auth/internal/controller"
	"otp-auth/internal/middleware"
	"otp-auth/internal/repositories"
	"otp-auth/internal/services"
	"otp-auth/internal/utils"
)

func main() {

	// Load configuration
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// initialize logger
	utils.InitLogger()

	// initialize redis
	utils.InitRedis(config.Redis)

	// Create repositories and services
	cacheRepo := repositories.NewCacheRepository()
	otpProvider := services.NewSmsOtpProvider()
	otpService := services.NewOtpService(cacheRepo, otpProvider)

	// Create controllers
	authController := controller.NewAuthController(otpService)

	// Set up Gin router
	router := gin.Default()

	// Register the Request ID middleware globally
	router.Use(middleware.StandardRequestMiddleware())

	// Define routes
	router.POST("/request-otp", authController.RequestOtp)
	router.POST("/verify-otp", authController.VerifyOTP)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("fail to start server")
		return
	}
}

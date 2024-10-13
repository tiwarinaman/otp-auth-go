package interfaces

type OtpProvider interface {
	GenerateOtp() string
	SendOtp(phoneNumber string, otp string) error
}

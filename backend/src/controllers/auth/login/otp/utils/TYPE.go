package loginUtils

import "github.com/golang-jwt/jwt/v5"

type Request struct {
	Email string `json:"email" validate:"required,email"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type VerifyRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type VerifyResponse struct {
	Success  bool             `json:"success"`
	Message  string           `json:"message"`
	Token    string           `json:"token,omitempty"`
	UserInfo *UserViewLabInfo `json:"user_info,omitempty"`
}

type GenerateJWTClaims struct {
	UserViewLabID int64  `json:"user_view_lab_id"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	jwt.RegisteredClaims
}

type UserViewLabInfo struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role,omitempty"`
}

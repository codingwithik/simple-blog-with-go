package dtos

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type RegisterRequest struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

// type MessageResponse struct {
// 	Message string `form:"message" json:"message" binding:"required"`
// }

type MessageResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

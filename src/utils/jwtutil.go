package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/codingwithik/simple-blog-backend-with-go/src/config"
	"github.com/codingwithik/simple-blog-backend-with-go/src/dtos"
	"github.com/codingwithik/simple-blog-backend-with-go/src/models"
	"github.com/dgrijalva/jwt-go"
)

var (
	jwtKey            = []byte(config.GetConfig().JWTSecret)
	expiryTime        = config.GetConfig().ExpiryTime
	refreshExpiryTime = config.GetConfig().RefreshExpiryTime
)

func GenerateJWT(u *models.User) (*dtos.TokenResponse, error) {
	nowTime := time.Now()
	expConf, err := strconv.Atoi(expiryTime)
	refreshExpConf, err := strconv.Atoi(refreshExpiryTime)

	expirationTime := nowTime.Add(time.Duration(int64(expConf)) * time.Hour)
	refreshExpirationTime := nowTime.Add(time.Duration(int64(refreshExpConf)) * time.Hour)

	claims := &dtos.JWTClaim{
		ID:    u.ID,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			Subject:   u.ID,
			IssuedAt:  nowTime.Unix(),
			Issuer:    "Blog",
			ExpiresAt: expirationTime.Unix(),
		},
	}

	refreshClaims := &dtos.JWTClaim{
		ID:    u.ID,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			Subject:   u.ID,
			IssuedAt:  nowTime.Unix(),
			Issuer:    "Blog",
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	token, err := tokenClaim.SignedString(jwtKey)
	refreshToken, err := refreshTokenClaim.SignedString(jwtKey)

	return &dtos.TokenResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpiresAt:    expirationTime.Unix(),
	}, err
}

func ValidateToken(signedToken string) (*dtos.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&dtos.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*dtos.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, nil
}

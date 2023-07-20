package middlewares

import (
	"github.com/codingwithik/simple-blog-backend-with-go/src/repositories"
	"github.com/gin-gonic/gin"
)

type middleware struct {
	Jwt             gin.HandlerFunc
	RefreshTokenJwt gin.HandlerFunc
}

var registeredMiddleWare middleware

func InitializeMiddleWares() {

	var (
		userRepo = repositories.NewUserRepo()
	)

	registeredMiddleWare = middleware{
		Jwt:             Auth(userRepo),
		RefreshTokenJwt: Auth(userRepo),
	}
}

func MiddleWare() *middleware {
	return &registeredMiddleWare
}

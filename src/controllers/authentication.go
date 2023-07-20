package controllers

import (
	"net/http"

	"github.com/codingwithik/simple-blog-backend-with-go/src/dtos"
	"github.com/codingwithik/simple-blog-backend-with-go/src/middlewares"
	"github.com/codingwithik/simple-blog-backend-with-go/src/services"
	"github.com/codingwithik/simple-blog-backend-with-go/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IAuthController interface {
	Login(ctx *gin.Context)
	//SendOTP(ctx *gin.Context)
	//ForgotPassword(ctx *gin.Context)
	Register(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type authController struct {
	authService services.IAuthService
}

// NewAuthController instantiates Auth Controller
func NewAuthController() IAuthController {
	return &authController{
		authService: services.NewAuthService(),
	}
}

func (ctl *authController) RegisterRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	auth.POST("/login", ctl.Login)
	auth.POST("/refresh-token", middlewares.MiddleWare().RefreshTokenJwt, ctl.RefreshToken)
	auth.POST("/register", ctl.Register)
	//auth.POST("/send_otp", ctl.SendOTP)
	//auth.POST("/reset_password", ctl.ForgotPassword)

}

func (ctl *authController) Login(ctx *gin.Context) {

	var body dtos.LoginRequest
	requestIdentifier := uuid.NewString()

	ctx.Header(utils.RequestIdentifier, requestIdentifier)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &dtos.MessageResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	res, err := ctl.authService.Login(body)

	if err != nil {
		err.HandleRequestErr(ctx)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (ctl *authController) Register(ctx *gin.Context) {

	var body dtos.RegisterRequest
	requestIdentifier := uuid.NewString()

	ctx.Header(utils.RequestIdentifier, requestIdentifier)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &dtos.MessageResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	res, err := ctl.authService.Register(body)

	if err != nil {
		err.HandleRequestErr(ctx)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (ctl *authController) RefreshToken(ctx *gin.Context) {

	requestIdentifier := uuid.NewString()
	ctx.Header(utils.RequestIdentifier, requestIdentifier)

	user, err := middlewares.UserFromContext(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &dtos.MessageResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	res, rErr := ctl.authService.RefreshToken(user)

	if rErr != nil {
		rErr.HandleRequestErr(ctx)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

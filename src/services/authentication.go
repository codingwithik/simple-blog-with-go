package services

import (
	"errors"
	"strings"

	"github.com/codingwithik/simple-blog-backend-with-go/src/config"
	"github.com/codingwithik/simple-blog-backend-with-go/src/dtos"
	"github.com/codingwithik/simple-blog-backend-with-go/src/exceptions"
	"github.com/codingwithik/simple-blog-backend-with-go/src/models"
	"github.com/codingwithik/simple-blog-backend-with-go/src/repositories"
	"github.com/codingwithik/simple-blog-backend-with-go/src/utils"
	"gorm.io/gorm"
)

var (
	validation = utils.Validation{}
	stringUtil = utils.StringUtil{}
)

type IAuthService interface {
	//ForgotPassword(rq *dtos.ForgotPasswordRequest) *exceptions.RequestError
	Login(body dtos.LoginRequest) (*dtos.TokenResponse, *exceptions.RequestError)
	//SendOTP(identity string) *errors.RequestError
	Register(body dtos.RegisterRequest) (*dtos.MessageResponse, *exceptions.RequestError)
	//ParseToken(token string) (*dtos.Claims, error)
	RefreshToken(user *models.User) (*dtos.TokenResponse, *exceptions.RequestError)
	//EditUserProfile(request *types.EditUserProfileRequest, user *model.User) error
}

type authService struct {
	jwtSecret string
	userRepo  repositories.IUserRepository
}

// NewAuthService will instantiate AuthService
func NewAuthService() IAuthService {
	return &authService{
		jwtSecret: config.GetConfig().JWTSecret,
		userRepo:  repositories.NewUserRepo(),
	}
}

func (as *authService) Register(rq dtos.RegisterRequest) (*dtos.MessageResponse, *exceptions.RequestError) {

	rq.Email = strings.ToLower(rq.Email)

	if !validation.ValidateEmail(rq.Email) {
		return nil, &exceptions.RequestError{
			StatusCode: 400,
			Err:        errors.New("email is invalid"),
		}
	}

	exists, err := as.userRepo.DoesEmailExist(rq.Email)

	if err != nil {
		return nil, &exceptions.RequestError{
			StatusCode: 500,
			Err:        errors.New("an error occurred"),
		}
	}

	if exists {
		return nil, &exceptions.RequestError{
			StatusCode: 400,
			Err:        errors.New("oops it appears the email you entered already exists"),
		}
	}

	//hash password
	passwordHash, err := stringUtil.BcryptHash(rq.Password)

	if err != nil {
		return nil, &exceptions.RequestError{
			StatusCode: 500,
			Err:        err,
		}
	}

	u := models.User{
		Name:     rq.Name,
		Email:    rq.Email,
		Password: passwordHash,
	}

	if err = as.userRepo.Create(&u); err != nil {
		return nil, &exceptions.RequestError{
			StatusCode: 400,
			Err:        errors.New("an error occurred when creating account"),
		}
	}

	return &dtos.MessageResponse{
		Status:  true,
		Message: "Registration Successful",
	}, nil

}

func (as *authService) Login(rq dtos.LoginRequest) (*dtos.TokenResponse, *exceptions.RequestError) {

	rq.Email = strings.ToLower(rq.Email)
	user, err := as.userRepo.FindByEmail(rq.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &exceptions.RequestError{
				StatusCode: 400,
				Err:        errors.New("invalid credentials"),
			}
		}
		return nil, &exceptions.RequestError{
			StatusCode: 500,
			Err:        errors.New("something went wrong please try again later"),
		}
	}

	err = stringUtil.CompareHash(rq.Password, user.Password)

	if err != nil {
		return nil, &exceptions.RequestError{
			StatusCode: 401,
			Err:        errors.New("invalid credentials"),
		}
	}

	issueResponse, err := utils.GenerateJWT(user)

	if err != nil {
		return nil, &exceptions.RequestError{
			StatusCode: 500,
			Err:        errors.New("an unexpected error occurred, try again"),
		}

	}

	return issueResponse, nil
}

func (as *authService) RefreshToken(user *models.User) (*dtos.TokenResponse, *exceptions.RequestError) {

	res, err := utils.GenerateJWT(user)

	if err != nil {
		return nil, &exceptions.RequestError{
			StatusCode: 500,
			Err:        errors.New("an unexpected error occurred, try again"),
		}
	}

	return res, nil
}

package middlewares

import (
	"errors"

	"net/http"

	"github.com/codingwithik/simple-blog-backend-with-go/src/dtos"
	"github.com/codingwithik/simple-blog-backend-with-go/src/models"
	"github.com/codingwithik/simple-blog-backend-with-go/src/repositories"
	"github.com/codingwithik/simple-blog-backend-with-go/src/utils"
	"github.com/gin-gonic/gin"
)

const (
	requestIdentifier         = "REQ_ID"
	AuthUserContextKey string = "auth"
)

func UserFromContext(ctx *gin.Context) (*models.User, error) {

	u, ok := ctx.Get(AuthUserContextKey)

	if !ok {
		return nil, errors.New("no user in context")
	}

	user := u.(models.User)

	return &user, nil
}

func Auth(ur repositories.IUserRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		user, err := ur.FindById(claims.ID)

		if err != nil {
			//logger.Error("error while extracting token: %v", err)
			context.AbortWithStatusJSON(http.StatusUnauthorized, dtos.MessageResponse{
				Status:  false,
				Message: "Unauthorized",
			})
			return
		}

		context.Set(AuthUserContextKey, *user)
		context.Next()
	}
}

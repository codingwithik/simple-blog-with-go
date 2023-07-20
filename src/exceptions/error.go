package exceptions

import (
	"github.com/gin-gonic/gin"
)

type RequestError struct {
	StatusCode int
	Err        error
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func (r *RequestError) HandleRequestErr(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(r.StatusCode, &ErrorResponse{
		Status:  false,
		Message: r.Err.Error(),
	})
}

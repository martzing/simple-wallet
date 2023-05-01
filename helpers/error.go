package helpers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomError interface {
	GetMessage() string
	GetStatusCode() int
}

type Error struct {
	Err        error
	Message    *string
	StatusCode *int
}

type ValidateError struct {
	Err        error
	Message    *string
	StatusCode *int
}

func (e *Error) GetMessage() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return *e.Message
}

func (e *Error) GetStatusCode() int {
	if e.Err != nil {
		return 500
	}
	return *e.StatusCode
}

func (ve *ValidateError) GetMessage() string {
	if ve.Err != nil {
		return ve.Err.Error()
	}
	msg := strings.Split(*ve.Message, "\n")[0]
	msg = strings.TrimSpace(strings.Split(msg, ":")[2])
	return msg
}

func (ve *ValidateError) GetStatusCode() int {
	if ve.Err != nil {
		return 500
	}
	return *ve.StatusCode
}

func AbortError(c *gin.Context, err any) {

	e, ok := err.(CustomError)
	if ok {
		c.AbortWithStatusJSON(e.GetStatusCode(), gin.H{
			"message": e.GetMessage(),
		})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong, Please contact support",
		})
	}
}

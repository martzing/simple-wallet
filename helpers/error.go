package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomError interface {
	GetMessage() string
	GetStatusCode() int
}

type Error struct {
	Message    string
	StatusCode int
}

type ValidateError struct {
	Message    string
	StatusCode int
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) GetStatusCode() int {
	return e.StatusCode
}

func (ve *ValidateError) GetMessage() string {
	fmt.Println(ve.Message)
	msg := strings.Split(ve.Message, "\n")[0]
	msg = strings.TrimSpace(strings.Split(msg, ":")[2])
	return msg
}

func (ve *ValidateError) GetStatusCode() int {
	return ve.StatusCode
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

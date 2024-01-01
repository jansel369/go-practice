package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Err struct {
	Message    *string
	StatusCode int
	StatusName string
	MainError  *error
}

type Result[T any] struct {
	Data  *T
	Error *Err
}

func (e Err) Error() string {
	return fmt.Sprintf(
		"Error %s %s: %s",
		e.StatusCode, e.StatusName,
		*e.Message,
	)
}

func StatusBadRequest(message string) Err {
	e := Err{
		Message:    &message,
		StatusCode: http.StatusBadRequest,
		StatusName: "BadRequest",
		MainError:  nil,
	}

	return e
}

func StatusUnauthorized(message string) Err {
	e := Err{
		Message:    &message,
		StatusCode: http.StatusUnauthorized,
		StatusName: "Unauthorized",
		MainError:  nil,
	}

	return e
}

func StatusNotFound(message string) Err {
	e := Err{
		Message:    &message,
		StatusCode: http.StatusNotFound,
		StatusName: "NotFound",
		MainError:  nil,
	}

	return e
}

func StatusInternalServerError(message string, err error) Err {
	e := Err{
		Message:    &message,
		StatusCode: http.StatusInternalServerError,
		StatusName: "InternalServerError",
		MainError:  &err,
	}

	return e
}

func (e Err) AbortRequest(c *gin.Context) {
	log.Printf(
		" => Abort Error %d %s: %s (%s)",
		e.StatusCode,
		e.StatusName,
		*e.Message,
		*e.MainError,
	)

	c.AbortWithStatusJSON(
		e.StatusCode,
		gin.H{
			"statusName": e.StatusName,
			"statusCode": e.StatusCode,
			"message":    e.Message,
		},
	)
}

// func ErrorHandler(c *gin.Context) gin.HandlerFunc {
// 	c.Next()

// 	for _, err := range c.Errors {
// 		println("from error handler", err)
// 	}

// 	c.JSON(http.StatusInternalServerError, "")
// }

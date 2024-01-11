package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

// var validate *validator.Validate
var validate = validator.New()

func ValidateMiddleware(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := c.ShouldBindBodyWith(model, binding.JSON); err != nil {

			log.Printf("\n  val err 1: %s", err)

			StatusInternalServerError("_val0x1", err).AbortRequest(c)

			return
		}

		// log.Println(model)

		// Use validator to check the request data
		if err := validate.Struct(model); err != nil {
			if invalidationErr, ok := err.(*validator.InvalidValidationError); ok {
				log.Printf(" => Invalid validation: %s %s\n", invalidationErr, err)

				StatusInternalServerError("_val0x2", err).AbortRequest(c)

				return
			}

			validationErrors := err.(validator.ValidationErrors)
			errorMessages := make(map[string]string)

			for _, fieldErr := range validationErrors {
				errorMessages[fieldErr.Field()] = fmt.Sprintf(
					"Field '%s' failed validation with tag '%s'",
					fieldErr.Field(),
					fieldErr.Tag(),
				)
			}

			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"statusName": "StatusBadRequest",
					"statusCode": http.StatusBadRequest,
					"message":    "Validation Error",
					"data":       errorMessages,
				},
			)
			return
		}

		c.Next()
	}
}

func FromBody(model interface{}, c *gin.Context) *Err {
	if err := c.ShouldBindBodyWith(model, binding.JSON); err != nil {
		err := StatusInternalServerError("_rdm0x1", errors.New("JSON binding faild"))

		return &err
	}

	return nil
}

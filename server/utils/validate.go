package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func ValidateMiddleware(inputStruct interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(inputStruct); err != nil {

			StatusInternalServerError("vmpl0x1", err).AbortRequest(c)

			return
		}

		// Use validator to check the request data
		if err := validate.Struct(inputStruct); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errorMessages := make(map[string]string)

			for _, fieldErr := range validationErrors {
				errorMessages[fieldErr.Field()] = fmt.Sprintf("Field '%s' failed validation with tag '%s'", fieldErr.Field(), fieldErr.Tag())
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

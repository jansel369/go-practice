package service

import (
	"log"
	"net/http"
	model "server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup, appCtx *utils.AppCtx) {

	router.POST(
		"/register",
		utils.ValidateMiddleware(&model.UserRegisterInput{}),
		func(c *gin.Context,
		) {
			log.Printf("test here register")
			var registerInput model.UserRegisterInput

			readErr := c.BindJSON(&registerInput)

			if readErr != nil {
				utils.StatusInternalServerError("rgiuxr1", readErr).AbortRequest(c)
				return
			}

			result := RegisterUser(registerInput, appCtx)

			if result.Data == nil {
				result.Error.AbortRequest(c)
				return
			}

			c.JSON(http.StatusOK, result.Data.Public())
		},
	)

	// router.POST(
	// 	"/login",
	// 	utils.ValidateMiddleware(&model.UserLoginInput{}),
	// 	func(ctx *gin.Context) {
	// 		user, err := model.FromJsonUser(c)

	// 		if err != nil {
	// 			println("error ", err)
	// 		}

	// 		c.JSON(http.StatusOK, user.Public())
	// 	},
	// )
}

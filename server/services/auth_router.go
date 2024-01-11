package service

import (
	"net/http"
	model "server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup, appCtx *utils.AppCtx) {

	router.POST(
		"/login",
		utils.ValidateMiddleware(&model.UserLoginInput{}),
		func(c *gin.Context) {
			var loginInput model.UserLoginInput

			if err := utils.FromBody(&loginInput, c); err != nil {
				err.AbortRequest(c)
				return
			}

			user, err := LoginUser(loginInput, appCtx)

			if err != nil {
				err.AbortRequest(c)
				return
			}

			c.JSON(http.StatusOK, user.Public())
		},
	)

	router.POST(
		"/register",
		utils.ValidateMiddleware(&model.UserRegisterInput{}),
		func(c *gin.Context) {
			var registerInput model.UserRegisterInput
			if !registerInput.FromBody(c) {
				return
			}

			user, error := RegisterUser(registerInput, appCtx)

			if error != nil {
				error.AbortRequest(c)
				return
			}

			c.JSON(http.StatusOK, user.Public())
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

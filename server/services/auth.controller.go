package service

import (
	model "server/models"
	"server/utils"
)

func RegisterUser(
	user model.UserRegisterInput,
	ctx *utils.AppCtx,
) utils.Result[model.User] {
	var existingUser model.User

	result := ctx.ORM.Where(
		"LOWER(email) = LOWER(?)",
		user.Email,
	).First(&existingUser)

	if result.Error != nil {
		e := utils.StatusInternalServerError("rgu0x1", result.Error)

		return utils.Result[model.User]{
			Data:  nil,
			Error: &e,
		}
	}

	if result.RowsAffected == 0 {
		e := utils.StatusBadRequest("User already exist")

		return utils.Result[model.User]{
			Data:  nil,
			Error: &e,
		}
	}

	hashedPass, hashErr := utils.HashPassword(user.Password)

	if hashErr != nil {
		e := utils.StatusInternalServerError("rgu0x2", hashErr)

		return utils.Result[model.User]{
			Data:  nil,
			Error: &e,
		}
	}

	createUser := model.User{
		Name:     user.Name,
		Age:      user.Age,
		Email:    user.Email,
		Password: hashedPass,
	}

	createResult := ctx.ORM.Create(&createUser)

	if createResult.Error != nil {
		e := utils.StatusInternalServerError("rgu0x3", createResult.Error)
		return utils.Result[model.User]{
			Data:  nil,
			Error: &e,
		}
	}

	return utils.Result[model.User]{
		Data: &createUser,
	}
}

package service

import (
	"log"
	model "server/models"
	"server/utils"
)

func RegisterUser(
	user model.UserRegisterInput,
	ctx *utils.AppCtx,
) (*model.User, *utils.Err) {
	var existingUser model.User

	result := ctx.ORM.Where(
		"LOWER(email) = LOWER(?)",
		user.Email,
	).First(&existingUser)

	if result.RowsAffected > 0 {
		e := utils.StatusBadRequest("User already exist")

		return nil, &e
	}

	if ormErr := utils.ORMError("rgu0x1", result); ormErr != nil {
		return nil, ormErr
	}

	log.Println(" => Users: ", existingUser)

	hashedPass, hashErr := utils.HashPassword(user.Password)

	if hashErr != nil {
		e := utils.StatusInternalServerError("rgu0x2", hashErr)

		return nil, &e
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

		return nil, &e
	}

	return &createUser, nil
}

func LoginUser(
	loginInput model.UserLoginInput,
	ctx *utils.AppCtx,
) (*model.User, *utils.Err) {
	var existingUser model.User

	result := ctx.ORM.Where(
		"LOWER(email) = LOWER(?)",
		loginInput.Email,
	).First(&existingUser)

	if result.RowsAffected == 0 {
		e := utils.StatusBadRequest("No user found")

		return nil, &e
	}

	if ormErr := utils.ORMError("rgu0x1", result); ormErr != nil {
		return nil, ormErr
	}

	if !utils.ComparePasswords(existingUser.Password, loginInput.Password) {
		e := utils.StatusBadRequest("Invalid email or password")

		return nil, &e
	}

	return &existingUser, nil
}

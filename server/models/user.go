package model

import (
	"time"

	_ "reflect"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        *uint  `gorm:primaryKey json:id`
	Name      string `json:name`
	Email     string `json:email`
	Password  string `json:password`
	Age       int    `json:age`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegisterInput struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,lte=8"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"gte=18"`
}

type UserLoginInput struct {
	Name     string `json:name validate:required`
	Password string `json:password validate:"required,lte=8"`
}

func (User) TableName() string {
	return "users"
}

func FromJsonUser(c *gin.Context) (User, error) {
	var user User

	err := c.BindJSON(&user)

	return user, err
}

func (u User) Public() map[string]any {
	return map[string]any{
		"name":      u.Name,
		"email":     u.Email,
		"password":  u.Password,
		"age":       u.Age,
		"updatedAt": u.UpdatedAt.UTC(),
		"createdAt": u.CreatedAt.UTC(),
	}
}

func (u *User) SetPassword(hashedPassword string) {
	u.Password = hashedPassword
}

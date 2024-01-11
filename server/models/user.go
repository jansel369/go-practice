package model

import (
	"log"
	"server/utils"
	"time"

	_ "reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true" sql:"AUTO_INCREMENT" json:id`
	Name      string    `json:name`
	Email     string    `json:email`
	Password  string    `json:password`
	Age       int       `json:age`
	CreatedAt time.Time `ogrm:"autoUpdateTime" json:createdAt`
	UpdatedAt time.Time `json:updatedAt`
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

func (u *UserRegisterInput) FromBody(c *gin.Context) bool {
	err := c.ShouldBindBodyWith(u, binding.JSON)

	if err != nil {
		log.Printf("parse error: %s", err)

		utils.StatusInternalServerError("fb_URI0x1", err).AbortRequest(c)

		return false
	}

	return true
}

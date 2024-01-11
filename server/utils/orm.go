package utils

import (
	"errors"

	"gorm.io/gorm"
)

func IsNotFound(result *gorm.DB) bool {
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

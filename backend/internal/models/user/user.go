package user

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username string `json:"username" gorm:"unique" validate:"required" form:"username"`
		Password string `json:"-" validate:"required" form:"password"`
		Salt     string `json:"-"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

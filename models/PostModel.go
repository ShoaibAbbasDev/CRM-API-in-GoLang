package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gorm.io/gorm"
)

type Company struct {
  
	gorm.Model
	Name      string
	Employees []Employee
}

type Employee struct {

	gorm.Model
	First_name  string
	Last_name   string
	Email       string
	Phone       string
	Probability int
	Stage       int
	CompanyID   uint
}

// company model validation

func (c Company) Validate() error {

	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(5, 50)),
	)
}

// employee model validation

func (e Employee) Validate() error {

	return validation.ValidateStruct(&e,
		validation.Field(&e.First_name, validation.Required, validation.Length(5, 50)),
		validation.Field(&e.Last_name, validation.Required, validation.Length(5, 50)),
		validation.Field(&e.Email, validation.Required, validation.Length(10, 50), is.Email),
		validation.Field(&e.Phone, validation.Required, validation.Length(5, 50)),
		validation.Field(&e.Probability, validation.Required),
		validation.Field(&e.Stage, validation.Required),
	)

}

package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
)

type Address struct {
	Street string
	City   string
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
	)
}

package main

import (
  "log"
  
  "gopkg.in/go-playground/validator.v9"
)

type Entity struct {
    Field string `validate:"awesome"` // <-- custom rule
}

func main() {
  v := validator.New()
  v.RegisterValidation("awesome", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "awesome"
	})
}

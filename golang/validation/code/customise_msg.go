package main

import (
  "log"

  "gopkg.in/go-playground/validator.v9"
)

type Entity struct {
    Field string `validate:"required"`
}

func main() {
  _ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}

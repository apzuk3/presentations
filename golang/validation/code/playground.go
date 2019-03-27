package main

import (
  "log"

  "gopkg.in/go-playground/validator.v9"
)

type CreateRequest struct {
    Title string `validate:"required,min=3,max=120"`
    Body  string `validate:"required,min=50,max=500"`
}

func main() {
  r := CreateRequest{}
  if err := validator.New().Struct(r); err != nil {
    log.Fatal(err)
  }
}

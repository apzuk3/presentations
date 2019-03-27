package main

import (
  "log"

  "gopkg.in/go-playground/validator.v9"
)

type CreateRequest struct {
    Title string `json:"title" validate:"required,min=3,max=120"`
    Body  string `json:"body" validate:"required,min=50,max=500"`
}

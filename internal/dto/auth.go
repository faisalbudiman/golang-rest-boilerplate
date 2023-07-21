package dto

import (
	"golang-rest-boilerplate/internal/model"
	res "golang-rest-boilerplate/pkg/util/response"
)

// Login
type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	model.UserEntityModel
}
type AuthLoginResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AuthLoginResponse `json:"data"`
	} `json:"body"`
}

// Register
type AuthRegisterRequest struct {
	model.UserEntity
}
type AuthRegisterResponse struct {
	Token string `json:"token"`
	model.UserEntityModel
}
type AuthRegisterResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data AuthRegisterResponse `json:"data"`
	} `json:"body"`
}

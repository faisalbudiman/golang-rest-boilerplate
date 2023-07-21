package dto

import (
	"golang-rest-boilerplate/internal/abstraction"
	"golang-rest-boilerplate/internal/model"
	res "golang-rest-boilerplate/pkg/util/response"
)

// Get
type UserGetRequest struct {
	abstraction.Pagination
	model.UserFilterModel
}
type UserGetResponse struct {
	Datas          []*model.UserEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type UserGetResponseDoc struct {
	Body struct {
		Meta res.Meta                `json:"meta"`
		Data []model.UserEntityModel `json:"data"`
	} `json:"body"`
}

// GetByCode
type UserGetByCodeRequest struct {
	Code string `param:"code" validate:"required"`
}
type UserGetByCodeResponse struct {
	model.UserEntityModel
}
type UserGetByCodeResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data UserGetByCodeResponse `json:"data"`
	} `json:"body"`
}

// Create
type UserCreateRequest struct {
	model.UserEntity
}
type UserCreateResponse struct {
	model.UserEntityModel
}
type UserCreateResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data UserCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type UserUpdateRequest struct {
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	UID   string `json:"uid"`
}
type UserUpdateResponse struct {
	model.UserEntityModel
}
type UserUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data UserUpdateResponse `json:"data"`
	} `json:"body"`
}

// ChangePassword
type UserChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
type UserChangePasswordResponse struct {
	model.UserEntityModel
}
type UserChangePasswordResponseDoc struct {
	Body struct {
		Meta res.Meta                   `json:"meta"`
		Data UserChangePasswordResponse `json:"data"`
	} `json:"body"`
}

// Delete
type UserDeleteRequest struct {
	Code string `param:"code" validate:"required"`
}
type UserDeleteResponse struct {
	model.UserEntityModel
}
type UserDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data UserDeleteResponse `json:"data"`
	} `json:"body"`
}

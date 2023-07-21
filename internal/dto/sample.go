package dto

import (
	"golang-rest-boilerplate/internal/abstraction"
	"golang-rest-boilerplate/internal/model"
	res "golang-rest-boilerplate/pkg/util/response"
)

// Get
type SampleGetRequest struct {
	abstraction.Pagination
	model.SampleFilterModel
}
type SampleGetResponse struct {
	Datas          []model.SampleEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type SampleGetResponseDoc struct {
	Body struct {
		Meta res.Meta                  `json:"meta"`
		Data []model.SampleEntityModel `json:"data"`
	} `json:"body"`
}

// GetById
type SampleGetByIdRequest struct {
	Id int64 `param:"id" validate:"required"`
}
type SampleGetByIdResponse struct {
	model.SampleEntityModel
}
type SampleGetByIdResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data SampleGetByIdResponse `json:"data"`
	} `json:"body"`
}

// Create
type SampleCreateRequest struct {
	model.SampleEntity
}
type SampleCreateResponse struct {
	model.SampleEntityModel
}
type SampleCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SampleCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type SampleUpdateRequest struct {
	Code string `param:"code" validate:"required"`
	model.SampleEntity
}

type SampleUpdateResponse struct {
	model.SampleEntityModel
}
type SampleUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SampleUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type SampleDeleteRequest struct {
	Code string `param:"code" validate:"required"`
}
type SampleDeleteResponse struct {
	model.SampleEntityModel
}
type SampleDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SampleDeleteResponse `json:"data"`
	} `json:"body"`
}

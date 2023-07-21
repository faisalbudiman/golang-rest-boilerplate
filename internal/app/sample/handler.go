package sample

import (
	"golang-rest-boilerplate/internal/abstraction"
	"golang-rest-boilerplate/internal/dto"
	"golang-rest-boilerplate/internal/factory"
	res "golang-rest-boilerplate/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

// Get
// @Summary Get sample
// @Description Get sample
// @Tags area
// @Accept json
// @Produce json
// @Security BearerAuth
// @param request query dto.SampleGetRequest true "request query"
// @Success 200 {object} dto.SampleGetResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /sample [get]
func (h *handler) Get(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.SampleGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Find(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get datas success", &result.PaginationInfo).Send(c)
}

// Get By ID
// @Summary Get sample by code
// @Description Get sample by code
// @Tags area
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "id path"
// @Success 200 {object} dto.SampleGetByIdRequestDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /sample/{id} [get]
func (h *handler) GetById(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.SampleGetByIdRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.FindById(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Create godoc
// @Summary Create sample
// @Description Create sample
// @Tags sample
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param request body dto.SampleCreateRequest true "request body"
// @Success 200 {object} dto.SampleCreateResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /sample [post]
func (h *handler) Create(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.SampleCreateRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Create(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update godoc
// @Summary Update sample
// @Description Update sample
// @Tags sample
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param code path int true "code path"
// @Param request body dto.SampleUpdateRequest true "request body"
// @Success 200 {object} dto.SampleUpdateResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /sample/{code} [put]
func (h *handler) UpdateByCode(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.SampleUpdateRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err)
	}

	result, err := h.service.UpdateByCode(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete godoc
// @Summary Delete sample
// @Description Delete sample
// @Tags sample
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param code path int true "code path"
// @Success 200 {object}  dto.SampleDeleteResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /sample/{code} [delete]
func (h *handler) DeleteByCode(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.SampleDeleteRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.DeleteByCode(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

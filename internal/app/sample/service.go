package sample

import (
	"errors"

	"golang-rest-boilerplate/internal/abstraction"
	"golang-rest-boilerplate/internal/dto"
	"golang-rest-boilerplate/internal/factory"
	"golang-rest-boilerplate/internal/model"
	"golang-rest-boilerplate/internal/repository"
	res "golang-rest-boilerplate/pkg/util/response"
	"golang-rest-boilerplate/pkg/util/trxmanager"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.SampleGetRequest) (*dto.SampleGetResponse, error)
	FindById(ctx *abstraction.Context, payload *dto.SampleGetByIdRequest) (*dto.SampleGetByIdResponse, error)
	Create(ctx *abstraction.Context, payload *dto.SampleCreateRequest) (*dto.SampleCreateResponse, error)
	UpdateByCode(ctx *abstraction.Context, payload *dto.SampleUpdateRequest) (*dto.SampleUpdateResponse, error)
	DeleteByCode(ctx *abstraction.Context, payload *dto.SampleDeleteRequest) (*dto.SampleDeleteResponse, error)
}

type service struct {
	Repository repository.Sample
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.SampleRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.SampleGetRequest) (*dto.SampleGetResponse, error) {
	var result *dto.SampleGetResponse
	var datas *[]model.SampleEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.SampleFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.SampleGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindById(ctx *abstraction.Context, payload *dto.SampleGetByIdRequest) (*dto.SampleGetByIdResponse, error) {
	var result *dto.SampleGetByIdResponse

	data, err := s.Repository.FindById(ctx, payload.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.SampleGetByIdResponse{
		SampleEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.SampleCreateRequest) (*dto.SampleCreateResponse, error) {
	var (
		result *dto.SampleCreateResponse
		req    model.SampleEntityModel
		data   *model.SampleEntityModel
	)

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		req.Context = ctx

		payload.Code = uuid.New()

		req.CreatedBy = ctx.Auth.ID
		req.SampleEntity = payload.SampleEntity
		data, err = s.Repository.Create(ctx, &req)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.SampleCreateResponse{
		SampleEntityModel: *data,
	}

	return result, nil
}

func (s *service) UpdateByCode(ctx *abstraction.Context, payload *dto.SampleUpdateRequest) (*dto.SampleUpdateResponse, error) {
	var (
		result *dto.SampleUpdateResponse
		req    model.SampleEntityModel
		data   *model.SampleEntityModel
	)

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByCode(ctx, payload.Code)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		req.Context = ctx
		req.SampleEntity = payload.SampleEntity
		data, err = s.Repository.Update(ctx, payload.Code, &req)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.SampleUpdateResponse{
		SampleEntityModel: *data,
	}

	return result, nil
}

func (s *service) DeleteByCode(ctx *abstraction.Context, payload *dto.SampleDeleteRequest) (*dto.SampleDeleteResponse, error) {
	var result *dto.SampleDeleteResponse
	var data *model.SampleEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.FindByCode(ctx, payload.Code)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data.Context = ctx
		data, err = s.Repository.Delete(ctx, payload.Code, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.SampleDeleteResponse{
		SampleEntityModel: *data,
	}

	return result, nil
}

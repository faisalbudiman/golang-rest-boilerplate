package auth

import (
	"golang-rest-boilerplate/internal/abstraction"
	"golang-rest-boilerplate/internal/dto"
	"golang-rest-boilerplate/internal/factory"
	"golang-rest-boilerplate/internal/model"
	"golang-rest-boilerplate/internal/repository"
	res "golang-rest-boilerplate/pkg/util/response"
	"golang-rest-boilerplate/pkg/util/trxmanager"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
}

type service struct {
	Repository repository.User
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.UserRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if data == nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(payload.Password)); err != nil {
		return result, res.CustomErrorBuilder(401, "unauthorized", "email or password doesn't match")
	}

	token, err := data.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err)
	}

	result = &dto.AuthLoginResponse{
		Token:           token,
		UserEntityModel: *data,
	}

	return result, nil
}

func (s *service) Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var (
		result   *dto.AuthRegisterResponse
		dataUser *model.UserEntityModel
	)

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		payload.IsActive = true
		dataUser, err = s.Repository.Create(ctx, &payload.UserEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	token, err := dataUser.GenerateToken()
	if err != nil {
		return nil, err
	}

	result = &dto.AuthRegisterResponse{
		Token:           token,
		UserEntityModel: *dataUser,
	}

	return result, nil
}

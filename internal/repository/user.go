package repository

import (
	"golang-rest-boilerplate/internal/abstraction"
	"golang-rest-boilerplate/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error)
	Create(ctx *abstraction.Context, data *model.UserEntity) (*model.UserEntityModel, error)
	FindById(ctx *abstraction.Context, id int) (*model.UserEntityModel, error)
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *user) FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Create(ctx *abstraction.Context, e *model.UserEntity) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UserEntityModel
	data.UserEntity = *e
	err := conn.Create(&data).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Update(ctx *abstraction.Context, id *int, e *model.UserEntityModel) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Model(e).Updates(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *user) FindById(ctx *abstraction.Context, id int) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) UpdateProfile(ctx *abstraction.Context, id int, e *model.UserEntityModel) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Model(e).Where("id = ?", id).Updates(map[string]interface{}{
		"email": e.Email,
		"name":  e.Name,
	}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	err = conn.Where("id = ?", id).First(&e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

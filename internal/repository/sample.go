package repository

import (
	"fmt"
	"golang-rest-boilerplate/internal/abstraction"
	"golang-rest-boilerplate/internal/model"

	"gorm.io/gorm"
)

type Sample interface {
	Find(ctx *abstraction.Context, m *model.SampleFilterModel, p *abstraction.Pagination) (*[]model.SampleEntityModel, *abstraction.PaginationInfo, error)
	FindById(ctx *abstraction.Context, id int64) (*model.SampleEntityModel, error)
	FindByCode(ctx *abstraction.Context, code string) (*model.SampleEntityModel, error)
	Create(ctx *abstraction.Context, e *model.SampleEntityModel) (*model.SampleEntityModel, error)
	Update(ctx *abstraction.Context, code string, e *model.SampleEntityModel) (*model.SampleEntityModel, error)
	Delete(ctx *abstraction.Context, code string, e *model.SampleEntityModel) (*model.SampleEntityModel, error)
}

type sample struct {
	abstraction.Repository
}

func NewSample(db *gorm.DB) *sample {
	return &sample{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *sample) Find(ctx *abstraction.Context, m *model.SampleFilterModel, p *abstraction.Pagination) (*[]model.SampleEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.SampleEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.SampleEntityModel{})

	// filter
	query = r.Filter(ctx, query, m)

	// sort
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}
	if p.SortBy == nil {
		sortBy := "created_at"
		p.SortBy = &sortBy
	}
	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	// pagination
	if p.Page == nil {
		page := 1
		p.Page = &page
	}
	if p.PageSize == nil {
		pageSize := 10
		p.PageSize = &pageSize
	}
	info = abstraction.PaginationInfo{
		Pagination: p,
	}
	limit := *p.PageSize + 1
	offset := (*p.Page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	err := query.Find(&datas).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, &info, err
	}

	info.Count = len(datas)
	info.MoreRecords = false
	if len(datas) > *p.PageSize {
		info.MoreRecords = true
		info.Count -= 1
		datas = datas[:len(datas)-1]
	}

	return &datas, &info, nil
}

func (r *sample) FindById(ctx *abstraction.Context, id int64) (*model.SampleEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.SampleEntityModel
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *sample) FindByCode(ctx *abstraction.Context, code string) (*model.SampleEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.SampleEntityModel
	err := conn.Where("code = ?", code).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *sample) Create(ctx *abstraction.Context, e *model.SampleEntityModel) (*model.SampleEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.First(&e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *sample) Update(ctx *abstraction.Context, code string, e *model.SampleEntityModel) (*model.SampleEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Model(e).Where("code = ?", code).Updates(map[string]interface{}{
		"name":        e.Name,
		"description": e.Description,
	}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	err = conn.Where("code = ?", code).First(&e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *sample) Delete(ctx *abstraction.Context, code string, e *model.SampleEntityModel) (*model.SampleEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("code = ?", code).First(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	err = conn.Model(e).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

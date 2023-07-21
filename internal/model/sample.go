package model

import (
	"golang-rest-boilerplate/internal/abstraction"
	"golang-rest-boilerplate/internal/type_constant"
	"golang-rest-boilerplate/pkg/util/date"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SampleEntity struct {
	Code        uuid.UUID                `json:"code" gorm:"index:unique;type:uuid;default:uuid_generate_v4()"`
	Name        string                   `json:"name" validate:"required" gorm:"index:idx_sample_name"`
	Description *string                  `json:"description"`
	SampleType  type_constant.SampleType `json:"sample_type"`
}

type SampleFilter struct {
	Code *string `query:"code" filter:"ILIKE"`
	Name *string `query:"name" filter:"ILIKE"`
}

type SampleEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	SampleEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SampleFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	SampleFilter
}

func (SampleEntityModel) TableName() string {
	return "sample"
}

func (m *SampleEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.ID
	return
}

func (m *SampleEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = date.DateTodayLocal()
	m.UpdatedBy = &m.Context.Auth.ID
	return
}

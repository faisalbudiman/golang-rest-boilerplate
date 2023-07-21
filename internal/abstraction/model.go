package abstraction

import (
	"time"

	"golang-rest-boilerplate/pkg/util/date"

	"gorm.io/gorm"
)

type Entity struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt time.Time  `json:"created_at"`
	CreatedBy int        `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *int       `json:"updated_by"`

	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Filter struct {
	CreatedAt *time.Time `query:"created_at"`
	CreatedBy *int       `query:"created_by"`
	UpdatedAt *time.Time `query:"updated_at"`
	UpdatedBy *int       `query:"updated_by"`
}

func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = date.DateTodayLocal()
	return
}

func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

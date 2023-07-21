package factory

import (
	"golang-rest-boilerplate/database"
	"golang-rest-boilerplate/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db               *gorm.DB
	SampleRepository repository.Sample
	UserRepository   repository.User
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("V0")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.SampleRepository = repository.NewSample(f.Db)
	f.UserRepository = repository.NewUser(f.Db)
}

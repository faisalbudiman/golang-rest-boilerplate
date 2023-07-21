package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	dbConnections map[string]*gorm.DB
)

func Init() {
	// configure database with PostgresSQL
	dbConfigurations := map[string]Db{
		"V0": &dbPostgreSQL{
			db: db{
				Host: os.Getenv("DB_HOST_V0"),
				User: os.Getenv("DB_USER_V0"),
				Pass: os.Getenv("DB_PASS_V0"),
				Port: os.Getenv("DB_PORT_V0"),
				Name: os.Getenv("DB_NAME_V0"),
			},
			SslMode: os.Getenv("DB_SSLMODE_V0"),
			Tz:      os.Getenv("DB_TZ_V0"),
		},
	}

	dbConnections = make(map[string]*gorm.DB)
	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection(name string) (*gorm.DB, error) {
	if dbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return dbConnections[strings.ToUpper(name)], nil
}

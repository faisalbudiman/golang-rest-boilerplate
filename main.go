package main

import (
	"os"

	db "golang-rest-boilerplate/database"
	"golang-rest-boilerplate/database/migration"
	"golang-rest-boilerplate/internal/factory"
	"golang-rest-boilerplate/internal/http"
	"golang-rest-boilerplate/internal/middleware"
	"golang-rest-boilerplate/pkg/util/env"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title golang-rest-boilerplate
// @version 0.0.1
// @description This is a doc for golang-rest-boilerplate.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /

func main() {
	var PORT = os.Getenv("PORT")

	// make sure database should the first init
	db.Init()

	// enabled migration, disabled if doesn't need
	migration.Init()

	// enable elasticsearch if needed
	// elasticsearch.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}

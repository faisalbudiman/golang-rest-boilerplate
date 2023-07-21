package http

import (
	"fmt"
	"net/http"
	"os"

	docs "golang-rest-boilerplate/docs"
	"golang-rest-boilerplate/internal/app/auth"
	"golang-rest-boilerplate/internal/app/sample"
	"golang-rest-boilerplate/internal/factory"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
		HOST    = os.Getenv("HOST")
		SCHEME  = os.Getenv("SCHEME")
	)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// doc
	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = HOST
	docs.SwaggerInfo.Schemes = []string{SCHEME}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	auth.NewHandler(f).Route(e.Group("/auth"))
	sample.NewHandler(f).Route(e.Group("/sample"))
}

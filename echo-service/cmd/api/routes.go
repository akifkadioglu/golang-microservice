package main

import (
	handlers "echoService/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func (app *Config) routes() {
	e := echo.New()
	e.GET("/", handlers.Greeting)
	e.Logger.Fatal(e.Start(":80"))
}

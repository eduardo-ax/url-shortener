package main

import (
	"github.com/eduardo-ax/url-shortener/api"
	"github.com/eduardo-ax/url-shortener/domain"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	v1g := e.Group("/v1")
	shortner := domain.NewBase62Shortener()
	h := api.NewUrlHandler(shortner)
	h.Register(v1g)
	e.Logger.Fatal(e.Start(":8080"))
}

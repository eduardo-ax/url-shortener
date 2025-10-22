package main

import (
	"github.com/eduardo-ax/url-shortener/api"
	"github.com/eduardo-ax/url-shortener/domain"
	"github.com/labstack/echo"
)

func main() {
	echoServer := echo.New()

	v1Group := echoServer.Group("/v1")

	shortener := domain.NewBase62Shortener()
	handler := api.NewUrlHandler(shortener)
	handler.Register(v1Group)

	echoServer.Logger.Fatal(echoServer.Start(":8080"))
}

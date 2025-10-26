package main

import (
	"github.com/eduardo-ax/url-shortener/api"
	"github.com/eduardo-ax/url-shortener/domain"
	"github.com/eduardo-ax/url-shortener/infrastructure"
	"github.com/labstack/echo"
)

func main() {
	echoServer := echo.New()
	v1Group := echoServer.Group("/v1")

	redis := infrastructure.NewCache()
	defer redis.Close()

	pool := infrastructure.NewPool()
	db := infrastructure.NewDatabase(pool)
	defer db.Close()

	shortener := domain.NewBase62Shortener(db, redis)
	handler := api.NewUrlHandler(shortener)
	handler.Register(v1Group)

	echoServer.Logger.Fatal(echoServer.Start(":8080"))
}

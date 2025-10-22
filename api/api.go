package api

import (
	"net/http"

	"github.com/eduardo-ax/url-shortener/domain"
	"github.com/labstack/echo"
)

type UrlRequest struct {
	URL string `json:"url"`
}

type UrlHandler struct {
	shortener domain.Shortener
}

func NewUrlHandler(shortener domain.Shortener) *UrlHandler {
	return &UrlHandler{
		shortener: shortener,
	}
}

func (u *UrlHandler) Register(e *echo.Group) {
	e.POST("/shorten", u.HandleShorten)
	e.GET("", u.HandleUrl)
}

func (u *UrlHandler) HandleShorten(c echo.Context) error {
	r := &UrlRequest{}
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format. Ensure the body is valid JSON.")
	}

	if r.URL == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "'url' field is required and cannot be empty.")
	}

	shortURL, err := u.shortener.Shorten(r.URL)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, shortURL)
}

func (u *UrlHandler) HandleUrl(c echo.Context) error {
	return nil
}

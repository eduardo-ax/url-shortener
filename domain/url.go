package domain

import (
	"errors"
	"net/url"
)

type URL struct {
	Url      string `json:"url"`
	ShortUrl string `json:"shortUrl"`
}

var ErrInvalidUrlFormat = errors.New("invalid URL format")

func NewUrl(longUrl string) (*URL, error) {

	if u, err := url.ParseRequestURI(longUrl); err != nil || u.Scheme == "" || u.Host == "" {
		return nil, ErrInvalidUrlFormat
	}

	return &URL{
		Url: longUrl,
	}, nil
}

type Shortener interface {
	Shorten(longUrl string) (URL, error)
}

type Base62Shortener struct{}

func NewBase62Shortener() *Base62Shortener {
	return &Base62Shortener{}
}

func (s *Base62Shortener) Shorten(longUrl string) (URL, error) {
	return URL{
		Url:      longUrl,
		ShortUrl: "https://short.com.br",
	}, nil
}

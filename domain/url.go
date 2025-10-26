package domain

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type URL struct {
	URL      string `json:"url"`
	ShortURL string `json:"shortUrl"`
}

var ErrInvalidUrlFormat = errors.New("invalid URL format")

func NewURL(longURL string) (URL, error) {
	if u, err := url.ParseRequestURI(longURL); err != nil || u.Scheme == "" || u.Host == "" {
		return URL{}, ErrInvalidUrlFormat
	}

	return URL{
		URL: longURL,
	}, nil
}

type Shortener interface {
	Shorten(ctx context.Context, longUrl string) (URL, error)
	Reverse(ctx context.Context, suffix string) (string, error)
}

type Base62Shortener struct {
	db    Storage
	redis Cache
}

func NewBase62Shortener(database Storage, cache Cache) *Base62Shortener {
	return &Base62Shortener{
		db:    database,
		redis: cache,
	}
}

func (s *Base62Shortener) Shorten(ctx context.Context, longUrl string) (URL, error) {
	originalURL, err := NewURL(longUrl)
	if err != nil {
		return URL{}, err
	}

	dbID, err := s.db.GetIdByUrl(ctx, originalURL)
	if err != nil {
		uniqueID, err := s.db.Persist(ctx, originalURL)
		if err != nil {
			return URL{}, err
		}
		dbID = uniqueID
	}
	base62, _ := convertToBase62(dbID)
	originalURL.ShortURL = "http://localhost:8080/v1/" + base62
	return originalURL, nil
}

func (s *Base62Shortener) Reverse(ctx context.Context, suffix string) (string, error) {
	cacheUrl, err := s.redis.Get(ctx, suffix)
	if err == nil && cacheUrl != "" {
		fmt.Println("URL ENCONTRADA NO CACHE!")
		return cacheUrl, nil
	}

	id, err := convertToDecimal(suffix)
	if err != nil {
		return "", err
	}

	url, err := s.db.GetById(ctx, id)
	if err != nil {
		return "", err
	}
	s.redis.Set(ctx, suffix, url)
	return url, nil
}

type Storage interface {
	GetIdByUrl(ctx context.Context, url URL) (int64, error)
	Persist(ctx context.Context, url URL) (int64, error)
	GetById(ctx context.Context, id int64) (string, error)
}

type Cache interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
}

const base62chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func convertToBase62(decimal int64) (string, error) {
	if decimal == 0 {
		return "0", nil
	}
	if decimal < 0 {
		return "", ErrInvalidUrlFormat
	}
	stringBase62 := ""
	numerator := decimal
	for numerator > 0 {
		remainder := numerator % 62
		stringBase62 = string(base62chars[remainder]) + stringBase62
		numerator = numerator / 62
	}
	return stringBase62, nil
}

func convertToDecimal(base62 string) (int64, error) {
	if !isValidBase62(base62) {
		return -1, ErrInvalidUrlFormat
	}

	if base62 == "0" {
		return 0, nil
	}

	var result int64
	for _, c := range base62 {
		idx := strings.IndexRune(base62chars, c)
		if idx < 0 {
			return -1, errors.New("invalid base62 encoding")
		}
		result = result*62 + int64(idx)
	}
	return result, nil
}

var base62Regex = regexp.MustCompile("^[a-zA-Z0-9]+$")

func isValidBase62(shortURL string) bool {
	if shortURL == "" {
		return false
	}
	return base62Regex.MatchString(shortURL)
}

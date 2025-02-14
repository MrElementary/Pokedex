package utils

import (
	"net/http"
	"time"

	"github.com/MrElementary/Pokedex/internal"
)

type Client struct {
	cache      internal.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: internal.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

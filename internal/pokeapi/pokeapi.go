package pokeapi

import (
	"net/http"
	//"encoding/json"
	"github.com/adamthiede/bootdev-pokedex/internal/pokecache"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func (client *Client) InvalidateCache() {
	client.cache.ReapNow()
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

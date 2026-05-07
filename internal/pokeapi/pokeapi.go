package pokeapi

import (
	"time"

	"github.com/jeanleeb/pokedexcli/internal/pokecache"
)

type Client struct {
	cache *pokecache.Cache
}

func NewClient() Client {
	return Client{
		cache: pokecache.NewCache(1 * time.Hour),
	}
}

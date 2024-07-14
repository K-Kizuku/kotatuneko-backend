package cache

import (
	"context"
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
)

type Client struct {
	con *ristretto.Cache
}

const ttl = 10 * time.Minute

func NewCacheClient() *Client {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 80000,  // number of keys to track frequency of (10M).
		MaxCost:     1 << 6, // maximum cost of cache (1GB).
		BufferItems: 64,     // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Client{
		con: cache,
	}
}

func (c *Client) Get(ctx context.Context, key string) (interface{}, bool) {
	val, founded := c.con.Get(key)
	return val, founded
}

func (c *Client) Set(ctx context.Context, key string, value interface{}) bool {
	added := c.con.SetWithTTL(key, value, 1, ttl)
	return added
}

func (c *Client) Delete(ctx context.Context, key string) {
	c.con.Del(key)
}

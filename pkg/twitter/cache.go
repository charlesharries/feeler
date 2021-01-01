package twitter

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/patrickmn/go-cache"
)

// CacheClient is a thin wrapper around cache.Cache allowing us to
// save the cache to a file on cache.Set().
type CacheClient struct {
	Cache    *cache.Cache
	Filepath string
}

// newCache creates a new cache client, saving data to the file
// at the given filepath. If there's already a cache at that
// filepath, it initialises the cache with the data found there.
func newCache(filepath string) *CacheClient {
	var items map[string]cache.Item

	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return newBlankCache(filepath)
	}

	err = json.Unmarshal(contents, &items)
	if err != nil {
		return newBlankCache(filepath)
	}

	return &CacheClient{
		Cache:    cache.NewFrom(5*time.Minute, 10*time.Minute, items),
		Filepath: filepath,
	}
}

// newBlankCache initialises a cache client at the given filepath.
func newBlankCache(filepath string) *CacheClient {
	return &CacheClient{
		Cache:    cache.New(5*time.Minute, 10*time.Minute),
		Filepath: filepath,
	}
}

// Save writes the currents state of the cache to the cache file.
func (c *CacheClient) Save() error {
	items := c.Cache.Items()

	raw, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.Filepath, raw, 0644)

	return err
}

// Get fetches an item from the cache.
func (c *CacheClient) Get(key string) (interface{}, bool) {
	return c.Cache.Get(key)
}

// Set sets an item in the cache.
func (c *CacheClient) Set(key string, val interface{}) error {
	c.Cache.Set(key, val, cache.DefaultExpiration)

	return c.Save()
}

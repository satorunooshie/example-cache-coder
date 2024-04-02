package cache

import (
	"sync"
	"time"

	"github.com/satorunooshie/example-cache-coder/cache/key"
)

type memory struct {
	sync.RWMutex
	cache map[string]data
}

type data struct {
	value any
	ttl   time.Duration
}

var c = memory{cache: make(map[string]data)}

func get(key key.Keyer) (any, error) {
	c.RLock()
	defer c.RUnlock()
	if v, ok := c.cache[key.Key()]; ok {
		return v.value, nil
	}
	return nil, nil
}

func set(key key.Keyer, value any) error {
	c.Lock()
	c.cache[key.Key()] = data{value: value, ttl: key.TTL()}
	c.Unlock()
	return nil
}

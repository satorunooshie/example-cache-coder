package cache

import (
	"github.com/satorunooshie/example-cache-coder/cache/coder"
	"github.com/satorunooshie/example-cache-coder/cache/key"
)

type Cache[T any] struct {
	coder coder.Coder[T]
}

func NewCache[T any](coder coder.Coder[T]) (*Cache[T], error) {
	if err := coder.Validate(); err != nil {
		return nil, err
	}
	return &Cache[T]{coder}, nil
}

func (c *Cache[T]) Get(k key.Keyer) (*T, error) {
	v, err := get(k)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	vv, ok := v.([]byte)
	if !ok {
		return nil, nil
	}
	return c.coder.Decode(vv)
}

func (c *Cache[T]) Set(k key.Keyer, v T) error {
	vv, err := c.coder.Encode(v)
	if err != nil {
		return err
	}
	return set(k, vv)
}

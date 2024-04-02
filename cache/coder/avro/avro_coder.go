package avro

import (
	"errors"

	icoder "github.com/satorunooshie/example-cache-coder/cache/coder"

	"github.com/hamba/avro/v2"
)

type coder[T any] struct {
	schema avro.Schema
}

var _ icoder.Coder[struct{}] = (*coder[struct{}])(nil)

func NewCoder[T any](s string) *coder[T] {
	return &coder[T]{schema: avro.MustParse(s)}
}

func (c coder[T]) Validate() error {
	if c.schema == nil {
		return errors.New("schema is not set")
	}
	return nil
}

func (c coder[T]) Encode(v T) ([]byte, error) {
	return avro.Marshal(c.schema, v)
}

func (c coder[T]) Decode(b []byte) (*T, error) {
	var v T
	if err := avro.Unmarshal(c.schema, b, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

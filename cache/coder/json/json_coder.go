package json

import (
	"encoding/json"

	"github.com/satorunooshie/example-cache-coder/cache/coder"
)

type Coder[T any] struct{}

var _ coder.Coder[struct{}] = (*Coder[struct{}])(nil)

func (Coder[T]) Validate() error {
	return nil
}

func (c Coder[T]) Encode(v T) ([]byte, error) {
	return json.Marshal(v)
}

func (c Coder[T]) Decode(b []byte) (*T, error) {
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/satorunooshie/example-cache-coder/cache/coder"
)

type Coder[T any] struct{}

var _ coder.Coder[struct{}] = (*Coder[struct{}])(nil)

func (Coder[T]) Validate() error {
	return nil
}

func (c Coder[T]) Encode(v T) ([]byte, error) {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (c Coder[T]) Decode(b []byte) (*T, error) {
	var v T
	if err := gob.NewDecoder(bytes.NewBuffer(b)).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

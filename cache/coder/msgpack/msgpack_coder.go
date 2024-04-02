package msgpack

import (
	"github.com/satorunooshie/example-cache-coder/cache/coder"

	"github.com/ugorji/go/codec"
)

type Coder[T any] struct{}

var _ coder.Coder[struct{}] = (*Coder[struct{}])(nil)

var h = codec.MsgpackHandle{}

func (Coder[T]) Validate() error {
	return nil
}

func (c Coder[T]) Encode(v T) ([]byte, error) {
	b := make([]byte, 0)
	if err := codec.NewEncoderBytes(&b, &h).Encode(v); err != nil {
		return nil, err
	}
	return b, nil
}

func (c Coder[T]) Decode(b []byte) (*T, error) {
	var v T
	if err := codec.NewDecoderBytes(b, &h).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

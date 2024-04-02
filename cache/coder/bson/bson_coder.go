package bson

import (
	"errors"

	"github.com/satorunooshie/example-cache-coder/cache/coder"

	"go.mongodb.org/mongo-driver/bson"
)

type Coder[T any] struct{}

var _ coder.Coder[struct{}] = (*Coder[struct{}])(nil)

func (Coder[T]) Validate() error {
	// avoid `WriteXXX can only write while positioned on a Element or Value but is positioned on a TopLevel`.
	// https://www.mongodb.com/docs/drivers/go/current/faq/#how-can-i-fix-the--writenull-can-only-write-while-positioned-on-a-element-or-value-but-is-positioned-on-a-toplevel--error-
	var t T
	switch any(t).(type) {
	case int:
		return errors.New("int is not supported")
	case string:
		return errors.New("string is not supported")
	case bool:
		return errors.New("bool is not supported")
	case []T, []*T:
		return errors.New("slice is not supported")
	default:
		return nil
	}
}

func (c Coder[T]) Encode(v T) ([]byte, error) {
	return bson.Marshal(v)
}

func (c Coder[T]) Decode(b []byte) (*T, error) {
	var v T
	if err := bson.Unmarshal(b, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

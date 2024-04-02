package protobuf

import (
	"google.golang.org/protobuf/proto"
)

type Coder[T any, PT PtrProtoMessage[T]] struct{}

func (Coder[T, PT]) Validate() error {
	return nil
}

func (c Coder[T, PT]) Encode(v T) ([]byte, error) {
	// FIXME: The Type Parameter is used to satisfy the coder.Coder Interface, but the Lock Interface cannot be copied, so modification is required.
	ptr := PT(&v)
	return proto.Marshal(ptr)
}

type PtrProtoMessage[T any] interface {
	proto.Message
	*T
}

func (c Coder[T, PT]) Decode(b []byte) (*T, error) {
	var t T
	ptr := PT(&t)
	if err := proto.Unmarshal(b, ptr); err != nil {
		return nil, err
	}
	return ptr, nil
}

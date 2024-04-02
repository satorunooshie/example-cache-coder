package flatbuf

import flatbuffers "github.com/google/flatbuffers/go"

type Handler[T any] interface {
	Make(b *flatbuffers.Builder, v *T) ([]byte, error)
	Read([]byte) (*T, error)
}

type Coder[H Handler[T], T any] struct{}

func (c Coder[H, T]) Validate() error {
	return nil
}

func (c Coder[H, T]) Encode(v T) ([]byte, error) {
	var h H
	builder := flatbuffers.NewBuilder(0)
	return h.Make(builder, &v)
}

func (c Coder[H, T]) Decode(b []byte) (*T, error) {
	var h H
	return h.Read(b)
}

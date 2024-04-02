package coder

type Coder[T any] interface {
	Validator
	Encode(T) ([]byte, error)
	Decode([]byte) (*T, error)
}

type Validator interface {
	Validate() error
}

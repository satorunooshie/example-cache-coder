package key

import "time"

type Keyer interface {
	Key() string
	TTL() time.Duration
}

type Key struct {
	str string
	ttl time.Duration
}

func (k Key) Key() string {
	return k.str
}

func (k Key) TTL() time.Duration {
	return k.ttl
}

func (k Key) String() string {
	return k.str
}

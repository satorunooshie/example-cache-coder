package key

import (
	"strconv"
	"time"
)

func NewKeyUserName(id int64, name string) Key {
	return Key{
		str: "user:" + strconv.FormatInt(id, 10) + ":name:" + name,
		ttl: 1 * time.Hour,
	}
}

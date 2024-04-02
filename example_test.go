package example

import (
	"fmt"

	"github.com/satorunooshie/example-cache-coder/cache"
	"github.com/satorunooshie/example-cache-coder/cache/coder/avro"
	"github.com/satorunooshie/example-cache-coder/cache/coder/bson"
	"github.com/satorunooshie/example-cache-coder/cache/coder/flatbuf"
	"github.com/satorunooshie/example-cache-coder/cache/coder/json"
	"github.com/satorunooshie/example-cache-coder/cache/coder/msgpack"
	"github.com/satorunooshie/example-cache-coder/cache/coder/protobuf"
	"github.com/satorunooshie/example-cache-coder/cache/key"
	"github.com/satorunooshie/example-cache-coder/flatc"
	"github.com/satorunooshie/example-cache-coder/model"
	"github.com/satorunooshie/example-cache-coder/proto"
)

func Example_coderProtobuf() {
	user := proto.User{
		Id:    1,
		Name:  "protobuf",
		Email: "test@protobuf.com",
	}
	c, _ := cache.NewCache(protobuf.Coder[proto.User, *proto.User]{})
	k := key.NewKeyUserName(1, user.Name)
	got1, _ := c.Get(k)
	_ = c.Set(k, user)
	got2, _ := c.Get(k)
	// ignore unexported fields.
	fmt.Println(got1.String(), got2.String())
	// Output:
	// <nil> name:"protobuf"  id:1  email:"test@protobuf.com"
}

func Example_coderMsgpack() {
	user := model.User{
		ID:    1,
		Name:  "msgpack",
		Email: "test@msgpack.com",
	}
	c, _ := cache.NewCache(msgpack.Coder[model.User]{})
	k := key.NewKeyUserName(1, user.Name)
	got1, _ := c.Get(k)
	_ = c.Set(k, user)
	got2, _ := c.Get(k)
	fmt.Println(got1, got2)
	// Output:
	// <nil> &{1 msgpack test@msgpack.com}
}

func Example_coderAvro() {
	user := model.User{
		ID:    1,
		Name:  "avro",
		Email: "test@avro.com",
	}
	c, _ := cache.NewCache(avro.NewCoder[model.User](user.Schema()))
	k := key.NewKeyUserName(1, user.Name)
	got1, _ := c.Get(k)
	_ = c.Set(k, user)
	got2, _ := c.Get(k)
	fmt.Println(got1, got2)
	// Output:
	// <nil> &{1 avro test@avro.com}
}

func Example_coderFlatbuf() {
	user := model.User{
		ID:    1,
		Name:  "flatbuffers",
		Email: "test@flatbuffers.com",
	}
	c, _ := cache.NewCache(flatbuf.Coder[flatc.UserHandler, model.User]{})
	k := key.NewKeyUserName(1, user.Name)
	got1, _ := c.Get(k)
	_ = c.Set(k, user)
	got2, _ := c.Get(k)
	fmt.Println(got1, got2)
	// Output:
	// <nil> &{1 flatbuffers test@flatbuffers.com}
}

func Example_coderJson() {
	user := model.User{
		ID:    1,
		Name:  "json",
		Email: "test@json.com",
	}
	c, _ := cache.NewCache(json.Coder[model.User]{})
	k := key.NewKeyUserName(1, user.Name)
	got1, _ := c.Get(k)
	_ = c.Set(k, user)
	got2, _ := c.Get(k)
	fmt.Println(got1, got2)
	// Output:
	// <nil> &{1 json test@json.com}
}

func Example_coderBson() {
	user := model.User{
		ID:    1,
		Name:  "bson",
		Email: "test@bson.com",
	}
	c, _ := cache.NewCache(bson.Coder[model.User]{})
	k := key.NewKeyUserName(1, user.Name)
	got1, _ := c.Get(k)
	_ = c.Set(k, user)
	got2, _ := c.Get(k)
	fmt.Println(got1, got2)
	// Output:
	// <nil> &{1 bson test@bson.com}
}

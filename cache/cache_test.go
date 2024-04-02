package cache

import (
	"reflect"
	"testing"
	"time"

	"github.com/satorunooshie/example-cache-coder/cache/coder"
	"github.com/satorunooshie/example-cache-coder/cache/coder/avro"
	"github.com/satorunooshie/example-cache-coder/cache/coder/bson"
	"github.com/satorunooshie/example-cache-coder/cache/coder/gob"
	"github.com/satorunooshie/example-cache-coder/cache/coder/json"
	"github.com/satorunooshie/example-cache-coder/cache/coder/msgpack"
)

type fakeKey struct {
	key string
}

func (k fakeKey) Key() string {
	return k.key
}

func (fakeKey) TTL() time.Duration {
	return 0
}

func TestCache(t *testing.T) {
	t.Parallel()
	testCache_int(t)
	testCache_string(t)
	testCache_bool(t)
	testCache_struct(t)
	testCache_map_struct(t)
	testCache_slice_struct(t)
}

func testCache_int(t *testing.T) {
	t.Helper()
	tests := []struct {
		name    string
		coder   coder.Coder[int]
		v       int
		wantErr bool
	}{
		{
			name:  "json_int",
			coder: json.Coder[int]{},
			v:     1,
		},
		{
			name:  "avro_int",
			coder: avro.NewCoder[int]("int"),
			v:     1,
		},
		{
			name:    "bson_int",
			coder:   bson.Coder[int]{},
			v:       1,
			wantErr: true,
		},
		{
			name:  "gob_int",
			coder: gob.Coder[int]{},
			v:     1,
		},
		{
			name:  "msgpack_int",
			coder: msgpack.Coder[int]{},
			v:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.v
			c, err := NewCache(tt.coder)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			k := fakeKey{tt.name}
			if err := c.Set(k, v); err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			got, err := c.Get(k)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("got nil, want %v", v)
				return
			}
			if !reflect.DeepEqual(*got, v) {
				t.Errorf("got %+v, want %+v", *got, v)
			}
		})
	}
}

func testCache_string(t *testing.T) {
	t.Helper()
	tests := []struct {
		name    string
		coder   coder.Coder[string]
		v       string
		wantErr bool
	}{
		{
			name:  "json_string",
			coder: json.Coder[string]{},
			v:     "string",
		},
		{
			name:  "avro_string",
			coder: avro.NewCoder[string]("string"),
			v:     "string",
		},
		{
			name:    "bson_string",
			coder:   bson.Coder[string]{},
			v:       "string",
			wantErr: true,
		},
		{
			name:  "gob_string",
			coder: gob.Coder[string]{},
			v:     "string",
		},
		{
			name:  "msgpack_string",
			coder: msgpack.Coder[string]{},
			v:     "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.v
			c, err := NewCache(tt.coder)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			k := fakeKey{tt.name}
			if err := c.Set(k, v); err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			got, err := c.Get(k)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("got nil, want %v", v)
				return
			}
			if !reflect.DeepEqual(*got, v) {
				t.Errorf("got %+v, want %+v", *got, v)
			}
		})
	}
}

func testCache_bool(t *testing.T) {
	t.Helper()
	tests := []struct {
		name    string
		coder   coder.Coder[bool]
		v       bool
		wantErr bool
	}{
		{
			name:  "json_bool",
			coder: json.Coder[bool]{},
			v:     true,
		},
		{
			name:  "avro_bool",
			coder: avro.NewCoder[bool]("boolean"),
			v:     true,
		},
		{
			name:    "bson_bool",
			coder:   bson.Coder[bool]{},
			v:       true,
			wantErr: true,
		},
		{
			name:  "gob_bool",
			coder: gob.Coder[bool]{},
			v:     true,
		},
		{
			name:  "msgpack_bool",
			coder: msgpack.Coder[bool]{},
			v:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.v
			c, err := NewCache(tt.coder)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			k := fakeKey{tt.name}
			if err := c.Set(k, v); err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			got, err := c.Get(k)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("got nil, want %v", v)
				return
			}
			if !reflect.DeepEqual(*got, v) {
				t.Errorf("got %+v, want %+v", *got, v)
			}
		})
	}
}

func testCache_struct(t *testing.T) {
	t.Helper()
	type testStruct struct {
		A int
		B string
	}
	tests := []struct {
		name    string
		coder   coder.Coder[testStruct]
		v       testStruct
		wantErr bool
	}{
		{
			name:  "json_struct",
			coder: json.Coder[testStruct]{},
			v:     testStruct{A: 1, B: "string"},
		},
		{
			name: "avro_struct",
			coder: avro.NewCoder[testStruct](`{
      "type":"record",
      "name":"testStruct",
      "fields":[
       {
          "name":"A",
          "type":"int"
       },
       {
          "name":"B",
          "type":"string"
       }
      ]
   }`),
			v: testStruct{A: 1, B: "string"},
		},
		{
			name:  "bson_struct",
			coder: bson.Coder[testStruct]{},
			v:     testStruct{A: 1, B: "string"},
		},
		{
			name:  "gob_struct",
			coder: gob.Coder[testStruct]{},
			v:     testStruct{A: 1, B: "string"},
		},
		{
			name:  "msgpack_struct",
			coder: msgpack.Coder[testStruct]{},
			v:     testStruct{A: 1, B: "string"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.v
			c, err := NewCache(tt.coder)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			k := fakeKey{tt.name}
			if err := c.Set(k, v); err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			got, err := c.Get(k)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("got nil, want %v", v)
				return
			}
			if !reflect.DeepEqual(*got, v) {
				t.Errorf("got %+v, want %+v", *got, v)
			}
		})
	}
}

func testCache_map_struct(t *testing.T) {
	t.Helper()
	type testStruct struct {
		A int
		B string
	}
	tests := []struct {
		name    string
		coder   coder.Coder[map[string]testStruct]
		v       map[string]testStruct
		wantErr bool
	}{
		{
			name:  "json_map_struct",
			coder: json.Coder[map[string]testStruct]{},
			v: map[string]testStruct{
				"1": {A: 1, B: "hello"},
				"2": {A: 2, B: "world"},
			},
		},
		{
			name: "avro_map_struct",
			coder: avro.NewCoder[map[string]testStruct](`{
      "type":"map",
      "values":{
        "type":"record",
        "name":"testStruct",
        "fields":[
         {
            "name":"A",
            "type":"int"
         },
         {
            "name":"B",
            "type":"string"
         }
        ]
      }
    }`),
			v: map[string]testStruct{
				"1": {A: 1, B: "hello"},
				"2": {A: 2, B: "world"},
			},
		},
		{
			name:  "bson_map_struct",
			coder: bson.Coder[map[string]testStruct]{},
			v: map[string]testStruct{
				"1": {A: 1, B: "hello"},
				"2": {A: 2, B: "world"},
			},
		},
		{
			name:  "gob_map_struct",
			coder: gob.Coder[map[string]testStruct]{},
			v: map[string]testStruct{
				"1": {A: 1, B: "hello"},
				"2": {A: 2, B: "world"},
			},
		},
		{
			name:  "msgpack_map_struct",
			coder: msgpack.Coder[map[string]testStruct]{},
			v: map[string]testStruct{
				"1": {A: 1, B: "hello"},
				"2": {A: 2, B: "world"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.v
			c, err := NewCache(tt.coder)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			k := fakeKey{tt.name}
			if err := c.Set(k, v); err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			got, err := c.Get(k)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("got nil, want %v", v)
				return
			}
			if !reflect.DeepEqual(*got, v) {
				t.Errorf("got %+v, want %+v", *got, v)
			}
		})
	}
}

func testCache_slice_struct(t *testing.T) {
	t.Helper()
	type testStruct struct {
		A int
		B string
	}
	tests := []struct {
		name    string
		coder   coder.Coder[[]testStruct]
		v       []testStruct
		wantErr bool
	}{
		{
			name:  "json_slice_struct",
			coder: json.Coder[[]testStruct]{},
			v: []testStruct{
				{A: 1, B: "hello"},
				{A: 2, B: "world"},
			},
		},
		{
			name: "avro_slice_struct",
			coder: avro.NewCoder[[]testStruct](`{
      "type":"array",
      "items":{
        "type":"record",
        "name":"testStruct",
        "fields":[
         {
            "name":"A",
            "type":"int"
         },
         {
            "name":"B",
            "type":"string"
         }
        ]
      }
    }`),
			v: []testStruct{
				{A: 1, B: "hello"},
				{A: 2, B: "world"},
			},
		},
		{
			name:  "bson_slice_struct",
			coder: bson.Coder[[]testStruct]{},
			v: []testStruct{
				{A: 1, B: "hello"},
				{A: 2, B: "world"},
			},
			wantErr: true,
		},
		{
			name:  "gob_slice_struct",
			coder: gob.Coder[[]testStruct]{},
			v: []testStruct{
				{A: 1, B: "hello"},
				{A: 2, B: "world"},
			},
		},
		{
			name:  "msgpack_slice_struct",
			coder: msgpack.Coder[[]testStruct]{},
			v: []testStruct{
				{A: 1, B: "hello"},
				{A: 2, B: "world"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.v
			c, err := NewCache(tt.coder)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			k := fakeKey{tt.name}
			if err := c.Set(k, v); err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			got, err := c.Get(k)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("got nil, want %v", v)
				return
			}
			if !reflect.DeepEqual(*got, v) {
				t.Errorf("got %+v, want %+v", *got, v)
			}
		})
	}
}

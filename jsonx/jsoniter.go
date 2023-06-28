package jsonx

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

type (
	// JsonIterator implements Client interface
	JsonIterator struct {
		client jsoniter.API
	}
)

func NewJSONIterator() *JsonIterator {
	return &JsonIterator{
		client: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

func (c *JsonIterator) Marshal(v interface{}) ([]byte, error) {
	return c.client.Marshal(v)
}

func (c *JsonIterator) Unmarshal(data []byte, v interface{}) error {
	return c.client.Unmarshal(data, v)
}

func (c *JsonIterator) Encode(writer io.Writer, v interface{}) error {
	return c.client.NewEncoder(writer).Encode(v)
}

func (c *JsonIterator) Decode(reader io.Reader, v interface{}) error {
	return c.client.NewDecoder(reader).Decode(v)
}

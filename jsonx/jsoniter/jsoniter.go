package jsoniter

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/leguminosa/gdk/jsonx"
)

// JsoniterClient implements StandardClient interface
type JsoniterClient struct {
	client jsoniter.API
}

func NewClient() jsonx.StandardClient {
	return &JsoniterClient{
		client: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

func (c *JsoniterClient) Marshal(v interface{}) ([]byte, error) {
	return c.client.Marshal(v)
}

func (c *JsoniterClient) Unmarshal(data []byte, v interface{}) error {
	return c.client.Unmarshal(data, v)
}

func (c *JsoniterClient) Encode(writer io.Writer, v interface{}) error {
	return c.client.NewEncoder(writer).Encode(v)
}

func (c *JsoniterClient) Decode(reader io.Reader, v interface{}) error {
	return c.client.NewDecoder(reader).Decode(v)
}

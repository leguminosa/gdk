package jsonx

import (
	"io"
	"sync"
)

type (
	// Client encapsulates json functionality
	Client interface {
		Marshal(v interface{}) ([]byte, error)
		Unmarshal(data []byte, v interface{}) error
		Encode(writer io.Writer, v interface{}) error
		Decode(reader io.Reader, v interface{}) error
	}
)

var (
	GetClient = initClient

	client   Client
	initOnce sync.Once
)

func initClient() Client {
	initOnce.Do(func() {
		client = NewJSONIterator()
	})

	return client
}

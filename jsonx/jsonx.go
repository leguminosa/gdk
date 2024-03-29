package jsonx

import (
	"io"
)

//go:generate mockgen -source=jsonx/jsonx.go -destination=jsonx/jsonx.mock.go -package=jsonx

// StandardClient encapsulates basic json functionality
type StandardClient interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Encode(writer io.Writer, v interface{}) error
	Decode(reader io.Reader, v interface{}) error
}

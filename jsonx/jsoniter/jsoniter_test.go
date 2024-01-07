package jsoniter

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert.NotEmpty(t, NewClient())
}

type TestStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestJsoniterClient_Marshal(t *testing.T) {
	c := NewClient()
	tests := []struct {
		name    string
		v       interface{}
		want    []byte
		wantErr bool
	}{
		{
			name:    "error",
			v:       make(chan int),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "success",
			v:       &TestStruct{ID: 1, Name: "test"},
			want:    []byte("{\"id\":1,\"name\":\"test\"}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Marshal(tt.v)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestJsoniterClient_Unmarshal(t *testing.T) {
	c := NewClient()
	dummyChan := make(chan int)
	tests := []struct {
		name    string
		data    []byte
		v       interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "error",
			data:    []byte("{\"id\":1,\"name\":\"test\"}"),
			v:       dummyChan,
			want:    dummyChan,
			wantErr: true,
		},
		{
			name:    "success",
			data:    []byte("{\"id\":1,\"name\":\"test\"}"),
			v:       &TestStruct{},
			want:    &TestStruct{ID: 1, Name: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.Unmarshal(tt.data, tt.v)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, tt.v)
		})
	}
}

func TestJsoniterClient_Encode(t *testing.T) {
	c := NewClient()
	tests := []struct {
		name    string
		v       interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "error",
			v:       make(chan int),
			want:    "",
			wantErr: true,
		},
		{
			name:    "success",
			v:       &TestStruct{ID: 1, Name: "test"},
			want:    "{\"id\":1,\"name\":\"test\"}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			err := c.Encode(writer, tt.v)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, writer.String())
		})
	}
}

func TestJsoniterClient_Decode(t *testing.T) {
	c := NewClient()
	dummyChan := make(chan int)
	tests := []struct {
		name    string
		reader  io.Reader
		v       interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "error",
			reader:  bytes.NewBufferString("{\"id\":1,\"name\":\"test\"}"),
			v:       dummyChan,
			want:    dummyChan,
			wantErr: true,
		},
		{
			name:    "success",
			reader:  bytes.NewBufferString("{\"id\":1,\"name\":\"test\"}"),
			v:       &TestStruct{},
			want:    &TestStruct{ID: 1, Name: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.Decode(tt.reader, tt.v)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, tt.v)
		})
	}
}

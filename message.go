package rosgo

import (
	"bytes"
)

type MessageType interface {
	Text() string
	MD5Sum() string
	Name() string
	NewMessage() Message
}

type Message interface {
	Type() MessageType
	Marshal(buf *bytes.Buffer) ([]byte, error)
	Unmarshal(buf *bytes.Reader) error
}

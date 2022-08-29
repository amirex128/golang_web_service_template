package kv

import (
	"io"
)

type Serializable interface {
	Encode(io.Writer) error
	Decode(io.Reader) error
}

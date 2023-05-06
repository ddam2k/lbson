package lbson

import (
	"github.com/ddam2k/lbson/lbson/marshal"
	"github.com/ddam2k/lbson/lbson/unmarshal"
)

func Marshal(data interface{}) ([]byte, error) {
	return marshal.Marshal(data)
}

func Unmarshal(bytes []byte) (interface{}, int) {
	return unmarshal.Unmarshal(bytes)
}

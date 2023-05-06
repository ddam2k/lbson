package lbson_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ddam2k/lbson"
)

func TestPrimitive(t *testing.T) {
	var u8 = 100

	if u8b, err := lbson.Marshal(u8); err != nil {
		t.Fail()
	} else {
		fmt.Printf("u8 = %x\n", u8b)
		if !bytes.Equal(u8b, []byte{0x21, 0x64}) {
			t.Fail()
		}
	}

	// var u16 = 1000

	// if u16b, err := lbson.Marshal(u16); err != nil {
	// 	t.Fail()
	// } else {
	// 	fmt.Printf("u16 = %x\n", u16b)
	// 	if !bytes.Equal(u16b, []byte{0x22, 0x64}) {
	// 		t.Fail()
	// 	}
	// }

}

package unmarshal

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/ddam2k/lbson/lbson/types"
)

func GetCodeAndNumber(bytes []byte) (int, int, int) {
	code := int(bytes[0] >> 4 & 0x0f)
	lenBytes := bytes[0] & 0x0f
	if lenBytes > 0x09 {
		return code, 0, int(lenBytes)
	}
	switch lenBytes {
	case 1:
		return code, int(bytes[1]), 1
	case 2:
		return code, int(binary.LittleEndian.Uint16(bytes[1:3])), 2
	case 4:
		return code, int(binary.LittleEndian.Uint32(bytes[1:5])), 4
	case 8:
		return code, int(binary.LittleEndian.Uint64(bytes[1:9])), 8
	}
	return 0, 0, 0
}

func BJStringU(bytes []byte) (interface{}, int) {
	_, len, spent := GetCodeAndNumber(bytes)
	str := bytes[spent+1 : spent+1+len]
	// idx += len
	fmt.Printf("str=%s\n", str)
	return string(str), spent + 1 + len
}

func BJMapU(bytes []byte) (interface{}, int) {
	var spent int
	out := make(map[string]interface{})
	_, count, spent := GetCodeAndNumber(bytes)
	fmt.Printf("map len = %d, bytes = %d\n", count, spent)
	idx := spent + 1
	for i := 0; i < count; i++ {
		var key string
		code, len, spent := GetCodeAndNumber(bytes[idx:])
		fmt.Printf("map code : %x %s\n", code, types.CodeToString(code))
		idx += spent + 1
		if code == types.BJTSTRING {
			key = string(bytes[idx : idx+len])
			idx += len
			fmt.Printf("key=%s\n", key)
		}
		if value, spent := Unmarshal(bytes[idx:]); value != nil {
			out[key] = value
			idx += spent
		}
	}
	return out, idx
}

func BJSliceU(bytes []byte) (interface{}, int) {
	var spent int
	out := make([]interface{}, 0)
	_, count, spent := GetCodeAndNumber(bytes)
	fmt.Printf("slice len = %d, bytes = %d\n", count, spent)
	idx := spent + 1
	for i := 0; i < count; i++ {
		if value, spent := Unmarshal(bytes[idx:]); value != nil {
			out = append(out, value)
			idx += spent
		}
	}
	return out, idx
}

func BJNumberU(bytes []byte) (interface{}, int) {
	code, no, spent := GetCodeAndNumber(bytes)
	switch code {
	case types.BJTINT:
		switch spent {
		case 1:
			return int8(no), 2
		case 2:
			return int16(no), 3
		case 4:
			return int32(no), 5
		case 8:
			return int64(no), 9
		}
	case types.BJTUINT:
		switch spent {
		case 1:
			return uint8(no), 2
		case 2:
			return uint16(no), 3
		case 4:
			return uint32(no), 5
		case 8:
			return uint64(no), 9
		}
	case types.BJTFLOAT:
		switch spent {
		case 4:
			return math.Float32frombits(uint32(no)), 5
		case 8:
			return math.Float64frombits(uint64(no)), 9
		}
	}
	return no, spent + 1
}

func BJBoolU(bytes []byte) (interface{}, int) {
	_, _, spent := GetCodeAndNumber(bytes)
	if spent == 0x0A {
		return false, 1
	} else if spent == 0x0B {
		return true, 1
	}
	return nil, 0
}

func Unmarshal(bytes []byte) (interface{}, int) {
	code := int(bytes[0] >> 4 & 0x0f)
	fmt.Printf("code: %d %s\n", code, types.CodeToString(code))
	switch code {
	case types.BJTMAP:
		return BJMapU(bytes)
	case types.BJTSTRING:
		return BJStringU(bytes)
	case types.BJTFLOAT:
		return BJNumberU(bytes)
	case types.BJTINT:
		return BJNumberU(bytes)
	case types.BJTUINT:
		return BJNumberU(bytes)
	case types.BJTSLICE:
		return BJSliceU(bytes)
	case types.BJTBOOL:
		return BJBoolU(bytes)
	}
	return nil, 0
}

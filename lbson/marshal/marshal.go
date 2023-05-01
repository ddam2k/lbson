package marshal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"

	"ddam2k/lbson/lbson/types"

	"golang.org/x/exp/constraints"
)

func BJNumber[T constraints.Integer | constraints.Float](n T) ([]byte, error) {
	bytes := make([]byte, 9)
	if n == T(uint8(n)) {
		bytes[0] = types.BJTUINT<<4 | 0x01
		bytes[1] = uint8(n)
		fmt.Printf("uint8 %v\n", n)
		return bytes[0:2], nil
	}
	if n == T(uint16(n)) {
		bytes[0] = types.BJTUINT<<4 | 0x02
		binary.LittleEndian.PutUint16(bytes[1:], uint16(n))
		fmt.Printf("uint16 %v\n", n)
		return bytes[0:3], nil
	}
	if n == T(uint32(n)) {
		bytes[0] = types.BJTUINT<<4 | 0x04
		binary.LittleEndian.PutUint32(bytes[1:], uint32(n))
		fmt.Printf("uint32 %v\n", n)
		return bytes[0:5], nil
	}
	if n == T(uint64(n)) {
		bytes[0] = types.BJTUINT<<4 | 0x08
		binary.LittleEndian.PutUint64(bytes[1:], uint64(n))
		fmt.Printf("uint64 %v\n", n)
		return bytes[0:9], nil
	}
	if n == T(int8(n)) {
		bytes[0] = types.BJTINT<<4 | 0x01
		bytes[1] = byte(int8(n))
		fmt.Printf("int8 %v\n", n)
		return bytes[0:2], nil
	}
	if n == T(int16(n)) {
		bytes[0] = types.BJTINT<<4 | 0x02
		binary.LittleEndian.PutUint16(bytes[1:], uint16(n))
		fmt.Printf("int16 %v\n", n)
		return bytes[0:3], nil
	}
	if n == T(int32(n)) {
		bytes[0] = types.BJTINT<<4 | 0x04
		binary.LittleEndian.PutUint32(bytes[1:], uint32(n))
		fmt.Printf("int32 %v\n", n)
		return bytes[0:5], nil
	}
	if n == T(int64(n)) {
		bytes[0] = types.BJTINT<<4 | 0x08
		binary.LittleEndian.PutUint64(bytes[1:], uint64(n))
		fmt.Printf("int64 %v\n", n)
		return bytes[0:9], nil
	}
	if n == T(float32(n)) {
		f32 := math.Float32bits(float32(n))
		bytes[0] = types.BJTFLOAT<<4 | 0x04
		binary.LittleEndian.PutUint32(bytes[1:], uint32(f32))
		fmt.Printf("float32 %v\n", n)
		return bytes[0:5], nil
	}
	if n == T(float64(n)) {
		f64 := math.Float64bits(float64(n))
		bytes[0] = types.BJTFLOAT<<4 | 0x08
		binary.LittleEndian.PutUint64(bytes[1:], uint64(f64))
		fmt.Printf("float64 %v\n", n)
		return bytes[0:9], nil
	}

	return bytes, fmt.Errorf("proper type not found")
}

func BJString(str string) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	fmt.Printf("str %v\n", str)
	strlen := len(str)
	buf.Write(LengthToBytes(types.BJTSTRING, strlen))
	buf.Write([]byte(str))
	return buf.Bytes(), nil
}

func BJBool(b bool) ([]byte, error) {
	fmt.Printf("bool %v\n", b)
	if b {
		return []byte{types.BJTBOOL<<4 | 0x0B}, nil
	} else {
		return []byte{types.BJTBOOL<<4 | 0x0A}, nil
	}
}

func Marshal(data interface{}) ([]byte, error) {
	t := reflect.ValueOf(data)
	switch t.Kind() {
	case reflect.Map:
		return BJMap(t)
	case reflect.Slice:
		return BJSlice(t)
	case reflect.Float32, reflect.Float64:
		return BJNumber(t.Float())
	case reflect.Bool:
		return BJBool(t.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return BJNumber(t.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return BJNumber(t.Uint())
	case reflect.String:
		return BJString(t.String())
	}
	return []byte{}, nil
}

func BJMap(value reflect.Value) ([]byte, error) {
	maplen := value.Len()
	fmt.Printf("map len = %d\n", maplen)
	buf := bytes.NewBuffer([]byte{})
	mapit := value.MapRange()
	buf.Write(LengthToBytes(types.BJTMAP, maplen))

	for next := mapit.Next(); next; next = mapit.Next() {
		key := mapit.Key().String()
		value := mapit.Value()
		fmt.Printf("map %v = \n", key)
		keylen := len(key)
		buf.Write(LengthToBytes(types.BJTSTRING, keylen))
		buf.Write([]byte(key))
		if result, err := Marshal(value.Interface()); err == nil {
			buf.Write(result)
		}
	}
	return buf.Bytes(), nil
}

func BJSlice(value reflect.Value) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	slicelen := value.Len()
	fmt.Printf("slice len = %d\n", slicelen)
	buf.Write(LengthToBytes(types.BJTSLICE, slicelen))
	for i := 0; i < value.Len(); i++ {
		v := value.Index(i)
		fmt.Printf("slice %v=\n", i)
		if result, err := Marshal(v.Interface()); err == nil {
			buf.Write(result)
		}
	}
	return buf.Bytes(), nil
}

func LengthToBytes(code uint8, l int) []byte {
	bytes := make([]byte, 5)
	if l > 65535 {
		bytes[0] = uint8(code<<4 | 0x04)
		binary.LittleEndian.PutUint32(bytes[1:], uint32(l))
		return bytes[0:5]
	} else if l > 255 {
		bytes[0] = byte(code<<4 | 0x02)
		binary.LittleEndian.PutUint16(bytes[1:], uint16(l))
		return bytes[0:3]
	} else {
		bytes[0] = byte(code<<4 | 0x01)
		fmt.Printf("code = %d %x", code, bytes[0])
		bytes[1] = uint8(l)
		fmt.Println(bytes[0:2])
		return bytes[0:2]
	}
}

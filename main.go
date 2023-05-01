package main

import (
	"encoding/json"
	"fmt"

	"ddam2k/lbson/lbson"
)

// const (
// 	BJTINT8    = '1'
// 	BJTINT16   = '2'
// 	BJTINT32   = '4'
// 	BJTINT64   = '8'
// 	BJTUINT8   = '!'
// 	BJTUINT16  = '@'
// 	BJTUINT32  = '$'
// 	BJTUINT64  = '*'
// 	BJTFLOAT32 = 'f'
// 	BJTFLOAT64 = 'F'
// 	BJTSTRING1 = 's'
// 	BJTSTRING2 = 't'
// 	BJTSTRING4 = 'r'
// 	BJTBOOL    = 'b'
// 	BJTMAP1    = 'M'
// 	BJTMAP2    = 'A'
// 	BJTMAP4    = 'P'
// 	BJTSLICE1  = 'S'
// 	BJTSLICE2  = 'L'
// 	BJTSLICE4  = 'I'
// )

func main() {
	fmt.Println("test1")

	const a = "{\"abcd\":1.5, \"bcdef\":\"mgkim\", \"cdefg\":[134,223,343,445,523, 1.1], \"boolean\": true}"

	var data map[string]interface{}

	// var data = map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }

	fmt.Printf("a: %d\n", len(a))
	json.Unmarshal([]byte(a), &data)

	bytes, err := lbson.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------------------------")
	fmt.Println(bytes)
	fmt.Printf("result len = %d\n", len(bytes))

	out, _ := lbson.Unmarshal(bytes)
	fmt.Println(out)
	// fmt.Println(bytes)

	// fmt.Println(data)

	// bytes1, _ := bson.Marshal(&data)

	// fmt.Printf("len: %d\n", len(bytes1))
	// fmt.Println(bytes1)

	// // var data2 interface{}
	// var data2 map[string]interface{}

	// bson.Unmarshal(bytes1, &data2)

	// fmt.Println("--------------------")

	// fmt.Println(data2)
}

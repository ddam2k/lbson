package main

import (
	"encoding/json"
	"fmt"

	"ddam2k/lbson/lbson"
)

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

package lbson_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ddam2k/lbson"
)

func TestMarshalUnMarshal(t *testing.T) {
	const jsInput = "{\"abcd\":1.5,\"bcdef\":\"mgkim\",\"boolean\":true,\"cdefg\":[134,223,343,65537,12345678,1.1]}"

	var data map[string]interface{}

	json.Unmarshal([]byte(jsInput), &data)

	bytes, err := lbson.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	out, _ := lbson.Unmarshal(bytes)

	jsOutBytes, _ := json.Marshal(out)

	if string(jsOutBytes) != jsInput {
		t.Fail()
	}
}

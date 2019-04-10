package orderedmap

import (
	"bytes"
	"encoding/json"
	"testing"
)

func Test_JsonMarshall(t *testing.T) {
	expected := []byte(`{"firstName":"Julius","lastName":"Caesar","age":50}`)
	dict := Collection{
		{"firstName", "Julius"},
		{"lastName", "Caesar"},
		{"age", 50},
	}
	json, _ := json.Marshal(dict)

	if bytes.Compare(expected, json) != 0 {
		t.Errorf("Expected %s, but got %s", expected, json)
	}
}

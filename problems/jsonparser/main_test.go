package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseToken(t *testing.T) {
	var tcases = []struct {
		name  string
		inp   string
		want  []string
		isErr bool
	}{
		{"invalid empty", "", []string{}, true},
		{"valid empty", "{}", []string{"{", "}"}, false},
		{"simple key value", `{"key": "value"}`, []string{"{", "key", ":", "value", "}"}, false},
		{"complex key value", `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value"
			"key5": 101,
			"key6": [1, 2, 3, 4]
		  }`, []string{"{", "key1", ":", "true", "key2", ":", "false", "key3", ":",
			"null", "key4", ":", "value", "key5", ":", "101", "key6", ":", "[", "1", "2", "3", "4", "]", "}"}, false},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tokenize([]byte(tt.inp))
			if err != nil && !tt.isErr {
				t.Errorf("Got %v, Wanted %v %v\n", got, tt.want, tt.isErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, Want %v\n", got, tt.want)
			}
		})
	}
}

func TestJSONify(t *testing.T) {
	var tcases = []struct {
		name string
		inp  string
		want map[string]interface{}
	}{
		{"valid empty", "{}", map[string]interface{}{}},
		{"simple key value", `{"key": "value"}`, map[string]interface{}{"key": "value"}},
		{"complex key value", `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 101,
			"key6": [1, 2, 3, 4]
		  }`, map[string]interface{}{"key1": true, "key2": false, "key3": nil, "key4": "value", "key5": 101, "key6": []any{1, 2, 3, 4}}},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := tokenize([]byte(tt.inp))
			if err != nil {
				t.Errorf("Previous step failed with err %v", err)
			}
			fmt.Println(tokens)
			res, err := JSONify(tokens)
			if err != nil {
				t.Errorf("jsonify failed with err %v", err)
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Got %v Want %v", res, tt.want)
			}
		})
	}
}

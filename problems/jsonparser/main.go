package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

func conv(token string) any {
	if token == "true" {
		return true
	} else if token == "false" {
		return false
	} else if token == "null" {
		return nil
	} else if n, err := strconv.Atoi(token); err == nil {
		return n
	} else {
		return token
	}
}

func JSONify(tokens []string) (map[string]interface{}, error) {
	var res = map[string]interface{}{}

	var i = 0
	var prev string = tokens[i]
	fmt.Println(tokens)

	i += 1
	for i < len(tokens) {
		if tokens[i] == "}" {
			return res, nil
		} else if tokens[i] == ":" && tokens[i+2] != "{" {
			res[prev] = conv(tokens[i+1])
			i += 2
		} else if tokens[i] == ":" && tokens[i+1] == "[" {
			ary := []any{}
			i += 2
			for tokens[i] != "]" {
				ary = append(ary, conv(tokens[i]))
				i += 1
			}
			res[prev] = ary
			i += 1
		} else {
			prev = tokens[i]
			i += 1
		}
	}
	return map[string]interface{}{}, errors.New("Something went wrong")
}

func tokenize(data []byte) ([]string, error) {
	tokens := []string{}
	i := 0
	for i < len(data) {
		ch := data[i]
		if ch == byte('{') || ch == byte(':') || ch == byte('}') || ch == byte('[') || ch == byte(']') {
			tokens = append(tokens, string(ch))
			i += 1
		} else if ch == byte('"') {
			j := i + 1
			valid := false
			escaped := false
			for j < len(data) {
				if data[j] == byte('"') && !escaped {
					valid = true
					break
				} else if data[j] == byte('\\') {
					escaped = true
				} else {
					escaped = false
					j += 1
				}
			}
			if !valid {
				return []string{}, errors.New("Unable to find string")
			} else {
				s := data[i+1 : j]
				tokens = append(tokens, string(s))
			}
			i = j + 1
		} else if i+4 < len(data) && reflect.DeepEqual(data[i:i+4], []byte("true")) {
			tokens = append(tokens, string(data[i:i+4]))
			i += 4
		} else if i+5 < len(data) && reflect.DeepEqual(data[i:i+5], []byte("false")) {
			tokens = append(tokens, string(data[i:i+5]))
			i += 5
		} else if i+4 < len(data) && reflect.DeepEqual(data[i:i+4], []byte("null")) {
			tokens = append(tokens, string(data[i:i+4]))
			i += 4
		} else if unicode.IsNumber(rune(ch)) {
			d := []byte{}
			for unicode.IsNumber(rune(data[i])) {
				d = append(d, data[i])
				i += 1
			}
			tokens = append(tokens, string(d))
		} else {
			i += 1
		}
	}
	if len(tokens) == 0 {
		return []string{}, errors.New("No tokens found")
	}
	return tokens, nil
}

func main() {
	s := "true"
	fmt.Printf("%v", reflect.DeepEqual([]byte(s), []byte("true")))
}

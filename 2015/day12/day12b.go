package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func sumA(f []interface{}) int {
	s := 0
	for _, u := range f {
		switch vv := u.(type) {
		case float64:
			s += int(vv)
		case string:
			s += 0
		case []interface{}:
			s += sumA(vv)
		default:
			s += sum(vv)
		}
	}
	return s
}

func sum(f interface{}) int {
	s := 0

	m := f.(map[string]interface{})
	for _, v := range m {
		switch vv := v.(type) {
		case string:
			if vv == "red" {
				return 0
			}
		case float64:
			s += int(vv)
		case []interface{}:
			s += sumA(vv)
		default:
			s += sum(vv)
		}
	}
	return s
}

func unmarshalAndSum(s string) int {
	var f interface{}
	err := json.Unmarshal([]byte(s), &f)
	if err != nil {
		panic("Failed to unmarshal input")
	}
	return sum(f)
}

func main() {
	fmt.Printf("%d\n", unmarshalAndSum("{\"e\": [1,2,3]}"))
	fmt.Printf("%d\n", unmarshalAndSum("{\"e\": [1,{\"c\":\"red\",\"b\":2},3]}"))
	fmt.Printf("%d\n", unmarshalAndSum("{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}"))
	fmt.Printf("%d\n", unmarshalAndSum("{\"e\": [1,\"red\",5]}"))

	t, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read input")
	}

	fmt.Printf("%d\n", unmarshalAndSum(string(t)))
}

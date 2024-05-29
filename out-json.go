package main

import (
	"encoding/json"
	"fmt"
)

type JSONData struct {
	Key   string  `json:"key"`
	Num   int     `json:"num"`
	Float float64 `json:"float"`
	Array []int   `json:"array"`
	Child struct {
		Key string `json:"key"`
	} `json:"child"`
}

func main() {
	data := JSONData{
		Key:   "value",
		Num:   1,
		Float: 1.1,
		Array: []int{1, 2, 3},
		Child: struct {
			Key string `json:"key"`
		}{
			Key: "value",
		},
	}

	v, _ := json.Marshal(data)

	fmt.Println(string(v))
}

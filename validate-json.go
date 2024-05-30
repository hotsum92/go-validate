package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
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
	var inputs []io.Reader

	if len(os.Args) == 1 {
		inputs = append(inputs, os.Stdin)
	} else {
		for _, arg := range os.Args[1:] {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", arg, err)
				os.Exit(1)
			}
			defer file.Close()
			inputs = append(inputs, file)
		}
	}

	var found bool

	for _, input := range inputs {
		var data JSONData
		decoder := json.NewDecoder(input)
		if err := decoder.Decode(&data); err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding JSON: %s\n", err)
			found = true
		}

		emptyFields := EmptyFields(data)

		if len(emptyFields) > 0 {
			fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", emptyFields)
			found = true
		}
	}

	if found {
		os.Exit(1)
	}
}

func IsZero(val interface{}) bool {
	var v reflect.Value

	if reflect.TypeOf(val).Kind() == reflect.Struct {
		v = reflect.ValueOf(val)
	}

	var typeName = v.Type().Name()

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name

		if v.Field(i).IsZero() {
			fmt.Printf("Error finding zero value: %s.%s\n", typeName, fieldName)
			return true
		}

		if v.Field(i).Kind() == reflect.Struct {
			return IsZero(v.Field(i).Interface())
		}
	}

	return false
}

func EmptyFields(s interface{}) []string {
	var emptyFields []string

	structVal := reflect.ValueOf(s)
	fieldNum := structVal.NumField()
	structType := reflect.TypeOf(s)

	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name

		isSet := field.IsValid() && !field.IsZero()

		if !isSet {
			emptyFields = append(emptyFields, fieldName)
		}

		if field.Kind() == reflect.Struct {
			emptyFields = append(emptyFields, EmptyFields(field.Interface())...)
		}
	}

	return emptyFields
}

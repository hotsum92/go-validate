package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

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
		r := csv.NewReader(input)

		for {
			row, err := r.Read()
			if err == io.EOF {
				break
			}

			fmt.Println(row)
		}
	}

	if found {
		os.Exit(1)
	}
}

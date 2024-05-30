package main

import (
	"fmt"
	"io"
	"os"
)

type RemoveBomReader struct {
	r io.Reader
}

func (r *RemoveBomReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if n > 0 && p[0] == 0xef {
		copy(p, p[3:])
		n -= 3
	}
	return
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

	r := io.MultiReader(inputs...)
	r = &RemoveBomReader{r}

	io.Copy(os.Stdout, r)
}

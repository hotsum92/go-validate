package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	var i int
	var s string = "123"
	i, _ = strconv.Atoi(s)
	fmt.Println(i) // -> 123

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")

	str := "abcde"
	fmt.Print(strings.Index(str, "fg"))
	fmt.Print(strings.Contains(str, "fg"))

	var reg = regexp.MustCompile(`\d+`)
	fmt.Println(reg.FindAllString("123 456 789", -1))
	fmt.Println(reg.MatchString("123 456 789"))
	fmt.Println(strings.TrimSpace(" 123 456 789 "))

	fmt.Println("%#v", t1)
	fmt.Printf("t1: %T\n", t1)
}

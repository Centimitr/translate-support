package main

import (
	"fmt"
	"github.com/Centimitr/translate-support/diff"
)

func main() {
	var o = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
	}
	var n = []string{
		"b",
		"f",
		"d",
		"f",
	}
	add, remove := diff.ModifiedLines(o, n)
	fmt.Println(add)
	fmt.Println(remove)
}

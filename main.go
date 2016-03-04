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
		"f",
		"d",
		"f",
	}
	result := diff.ModifiedLines(o, n)
	for _, l := range result.Lines {
		fmt.Println(l.Text)
	}
}

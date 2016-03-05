package main

import (
	"fmt"
	"github.com/Centimitr/translate-support/diff"
	"io/ioutil"
	// "strings"
)

func main() {
	oBytes, _ := ioutil.ReadFile("old.txt")
	nBytes, _ := ioutil.ReadFile("new.txt")
	fmt.Println(oBytes, nBytes)
	o := []string{
		"a",
		"b",
		"c",
		"d",
		"f",
	}
	n := []string{
		"x",
		"x",
		"b",
		"c",
		"x",
		"d",
		"x",
		"f",
		"x",
	}
	result := diff.DiffResult(o, n)
	fmt.Println(result.String())
}

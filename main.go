package main

import (
	"fmt"
	"github.com/Centimitr/translate-support/diff"
	"io/ioutil"
	"strings"
)

func main() {
	oBytes, _ := ioutil.ReadFile("old.txt")
	nBytes, _ := ioutil.ReadFile("new.txt")
	result := diff.ModifiedLines(strings.Split(string(oBytes), ("\n")), strings.Split(string(nBytes), ("\n")))
	for _, l := range result.Lines {
		switch {
		case l.IsUnmodified:
			fmt.Print("   ")
		case l.IsAdded:
			fmt.Print(" + ")
		case l.IsRemoved:
			fmt.Print(" - ")
		}
		fmt.Println(l.Text)
	}
}

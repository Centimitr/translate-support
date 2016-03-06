package main

import (
	"bufio"
	"fmt"
	"github.com/Centimitr/translate-support/diff"
	// "io/ioutil"
	// "strings"
	"os"
)

func ReadFileLines(filename string) (lines []string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return
}

func main() {
	result := diff.DiffResult(ReadFileLines("old.txt"), ReadFileLines("new.txt"))
	fmt.Println(result.String())
}

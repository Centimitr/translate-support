package main

import (
	"github.com/Centimitr/translate-support/diff"
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

func DiffLatest(filename string) string {
	//read watch and judge whether it is in the list
	prevVer := "alpha.0"
	latestVer := "alpha.1"
	//read old & new text
	oldText := ReadFileLines(WORKSPACE_DIR + prevVer + "/" + filename)
	newText := ReadFileLines(WORKSPACE_DIR + lastestVer + "/" + filename)
	//diff
	result := diff.DiffResult(oldText, newText)
	//marshal json
	return "[]"
}

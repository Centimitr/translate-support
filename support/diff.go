package support

import (
	"bufio"
	"fmt"
	"github.com/Centimitr/translate-support/diff"
	"os"
	"path/filepath"
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

func (c *Config) DiffLatest(filename string) string {
	//read watch and judge whether it is in the list
	prevVer := "alpha.0"
	latestVer := "alpha.1"
	//read old & new text
	oldText := ReadFileLines(filepath.Join(WORKSPACE_DIR, c.GetSrcLang(), prevVer, filename))
	newText := ReadFileLines(filepath.Join(WORKSPACE_DIR, c.GetSrcLang(), latestVer, filename))
	//diff
	diff.DiffResult(oldText, newText)
	//marshal json
	return "[]"
}

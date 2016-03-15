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

func (c *Config) LineDiff(curVer, oldVer, filename string) string {
	//read watch and judge whether it is in the list
	//read old & new text
	oldText := ReadFileLines(filepath.Join(WORKSPACE_DIR, c.GetSrcLang(), oldVer, filename))
	curText := ReadFileLines(filepath.Join(WORKSPACE_DIR, c.GetSrcLang(), curVer, filename))
	//diff
	diff.LineDiff(oldText, curText)
	//marshal json
	return "[]"
}

package main

import (
	"bufio"
	"fmt"
	// "github.com/Centimitr/translate-support/diff"
	spt "github.com/Centimitr/translate-support/support"
	// "io/ioutil"
	// "strings"
	"os"
)

func main() {
	// result := diff.DiffResult(ReadFileLines("old.txt"), ReadFileLines("new.txt"))
	// fmt.Println(result.String())
	spt.Init()
	spt.AddSrcLang("en-us")
	spt.AddTgtLang("zh-cn")
	spt.AddWatch("test.txt")
	spt.CopyFormerTrans()
	spt.DiffLatest("test.txt")
	spt.GenResult()
}

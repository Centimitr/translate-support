package main

import (
	// "github.com/Centimitr/translate-support/diff"
	spt "github.com/Centimitr/translate-support/support"
	// "io/ioutil"
	// "strings"
)

func main() {
	// result := diff.DiffResult(ReadFileLines("old.txt"), ReadFileLines("new.txt"))
	// fmt.Println(result.String())
	spt.Init()
	spt.AddSrcLang("en-us")
	spt.AddTgtLang("zh-cn")
	spt.AddWatch("test.txt")
	spt.CopyFormerTrans([]string{})
	spt.DiffLatest("test.txt")
	spt.GenResult()
}

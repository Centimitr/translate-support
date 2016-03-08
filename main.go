package main

import (
	// "github.com/Centimitr/translate-support/diff"
	"github.com/Centimitr/translate-support/support"
	// "io/ioutil"
	// "strings"
)

func main() {
	// result := diff.DiffResult(ReadFileLines("old.txt"), ReadFileLines("new.txt"))
	// fmt.Println(result.String())
	var spt support.Config
	spt.Init()
	spt.SetSrcLang("en-us")
	spt.SetTgtLang("zh-cn")
	spt.AddVer("alpha.0")
	spt.AddWatch("test.txt")
	spt.CopyFormerTrans([]string{})
	spt.CreateEmptyTrans([]string{})
	spt.DiffLatest("test.txt")
	spt.GenResult()
}

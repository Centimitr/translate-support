package main

import (
	// "github.com/Centimitr/translate-support/diff"
	"github.com/Centimitr/translate-support/support"
	// "io/ioutil"
	// "strings"
	"fmt"
	"html"
	"net/http"
	"os"
)

func main() {
	// result := diff.DiffResult(ReadFileLines("old.txt"), ReadFileLines("new.txt"))
	// fmt.Println(result.String())
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "init":
			support.Init()
		case "serve":
			// var spt = support.Ins()
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
			})
			http.ListenAndServe(":4567", nil)
		case "proc":
			var spt = support.Ins()
			// var spt = new(support.Config)
			spt.SetSrcLang("en-us")
			spt.SetTgtLang("zh-cn")
			spt.AddVer("alpha.0")
			spt.AddWatch("test.txt")
			spt.CopyFormerTrans([]string{})
			spt.CreateEmptyTrans([]string{})
			spt.LineDiffLatest("test.txt")
			spt.GenResult()
		case "test":
			var spt = support.Ins()
			spt.SetSrcLang("en-us")
			spt.SetTgtLang("zh-cn")
			spt.AddVer("alpha.0")
			spt.AddWatch("test.txt")
		}
	}
}

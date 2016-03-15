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
			if spt, initialed := support.Ins(); initialed {
				fmt.Println("initialed.")
				wd, _ := os.Getwd()
				http.Handle("/", http.FileServer(http.Dir(wd+"/"+support.WORKSPACE_DIR+"/dist")))
				http.HandleFunc("/api/versions", func(w http.ResponseWriter, r *http.Request) {
					// get post
					fmt.Println(r.Method)
					switch r.Method {
					case "GET":
						fmt.Println(spt.GetVers())
					case "POST":
						fmt.Println(spt.GetVers())
						spt.AddVer(r.Header.Get("name"))
						fmt.Println(spt.GetVers())
					}
				})
				http.HandleFunc("/api/{version}", func(w http.ResponseWriter, r *http.Request) {
					// delete put(basever) put(initwatch) put(inittranslate)
				})
				http.HandleFunc("/api/{version}/watch", func(w http.ResponseWriter, r *http.Request) {
					// get post delete
				})
				http.HandleFunc("/api/{version}/compare/{oldver}/", func(w http.ResponseWriter, r *http.Request) {
				})

				// test
				http.HandleFunc("/t", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
				})
				http.ListenAndServe(":4567", nil)
			}
			// case "proc":
			// 	var spt = support.Ins()
			// 	// var spt = new(support.Config)
			// 	spt.SetSrcLang("en-us")
			// 	spt.SetTgtLang("zh-cn")
			// 	spt.AddVer("alpha.0")
			// 	spt.AddWatch("test.txt")
			// 	spt.CopyFormerTrans([]string{})
			// 	spt.CreateEmptyTrans([]string{})
			// 	spt.LineDiffLatest("test.txt")
			// 	spt.GenResult()
			// case "test":
			// 	var spt = support.Ins()
			// 	spt.SetSrcLang("en-us")
			// 	spt.SetTgtLang("zh-cn")
			// 	spt.AddVer("alpha.0")
			// 	spt.AddWatch("test.txt")
		}
	}
}

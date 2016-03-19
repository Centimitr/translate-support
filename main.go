package main

import (
	// "github.com/Centimitr/translate-support/diff"
	"github.com/Centimitr/translate-support/support"
	"github.com/Centimitr/translate-support/tool"
	// "io/ioutil"
	// "encoding/json"
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
				fmt.Println("Initialed.")
				wd, _ := os.Getwd()
				http.Handle("/", http.FileServer(http.Dir(wd+"/"+support.WORKSPACE_DIR+"/dist")))
				http.HandleFunc("/api/versions", func(w http.ResponseWriter, r *http.Request) {
					// get post
					switch r.Method {
					case "GET":
						d := spt.GetVers()
						fmt.Fprint(w, d)
					case "POST":
						spt.AddVer(r.Header.Get("version"))
					}
				})
				// to avoid using of 3rd party mux library, router params transferred to get
				http.HandleFunc("/api/versions/v", func(w http.ResponseWriter, r *http.Request) {
					// delete put(basever) put(initwatch) put(inittranslate)
					switch r.Method {
					case "DELETE":
						spt.RemoveVer(r.Header.Get("version"))
					case "PUT":
						if r.Header.Get("initWatch") == "true" {

						}
						if r.Header.Get("initTranslate") == "true" {

						}

					}
				})
				http.HandleFunc("/api/versions/v/watches/", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Access-Control-Allow-Origin", "*")
					w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,HEAD,OPTIONS")
					// get post delete
					r.ParseForm()
					switch r.Method {
					case "GET":
						d := spt.GetWatchs(r.Form.Get("version"))
						fmt.Fprint(w, tool.StringSliceToJson(d))
					case "POST":
						spt.AddWatch(r.Form.Get("version"), r.Form.Get("filepath"))
						d := spt.GetWatchs(r.Form.Get("version"))
						fmt.Fprint(w, tool.StringSliceToJson(d))
					case "DELETE":
						spt.RemoveWatch(r.Form.Get("version"), r.Form.Get("filepath"))
						d := spt.GetWatchs(r.Form.Get("version"))
						fmt.Fprint(w, tool.StringSliceToJson(d))
					}
				})
				http.HandleFunc("/api/versions/v/compare/", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Access-Control-Allow-Origin", "*")
					switch r.Method {
					case "GET":
						r.ParseForm()
						d := spt.LineDiff(r.Form.Get("version"), r.Form.Get("oldVersion"), r.Form.Get("filepath"))
						fmt.Fprint(w, d)
					}
				})

				// test
				http.HandleFunc("/t", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
				})
				http.ListenAndServe(":4567", nil)
			}
		}
	}
}

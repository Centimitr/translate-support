package support

import (
	"os"
)

const WORKSPACE_DIR = "translate_workspace"

type Version struct {
	Name `json:name`
	// GitHash `json:gitHash`
}

type Config struct {
	SrcLang  string    `json:source`
	TgtLang  string    `json:target`
	Versions []Version `json:versions`
	Watch    []string  `json:watch`
}

var config Config

func Init() {
	os.Mkdir("./"+WORKSPACE_DIR, 0777)
	os.Create("./" + WORKSPACE_DIR + "/config.json")
}

func addLang(dirname string) error {
	return os.Mkdir("./"+WORKSPACE_DIR+"/"+dirname, 0777)
}

func AddSrcLang(dirname string) error {
	addLang(dirname)
}

func AddTgtLang(dirname string) error {
	addLang(dirname)
}

func AddWatch(filename string) error {

}

func ReadWatch() {

}

package main

import (
	"os"
)

const WORKSPACE_DIR = "workspace"

type Config struct {
	SrcLang string `json:source`
	TgtLang string `json:target`
}

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

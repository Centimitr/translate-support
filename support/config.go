package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	// "strconv"
	// "time"
)

const (
	WORKSPACE_DIR   = "translate_workspace"
	CONFIG_FILENAME = "config.json"
)

type Version struct {
	Name    string   `json:"name"`
	BaseVer string   `json:"basever"`
	Watch   []string `json:"watch"`
	Time    string   `json:"time"`
	// GitHash `json:gitHash`
}

type Config struct {
	Port     string    `json:"port"`
	SrcLang  string    `json:"source"`
	TgtLang  string    `json:"target"`
	Versions []Version `json:"versions"`
}

func Ins() (Config, bool) {
	var c Config
	d, e := ioutil.ReadFile(filepath.Join(WORKSPACE_DIR, CONFIG_FILENAME))
	if e != nil {
		return c, false
	}
	e = json.Unmarshal(d, &c)
	if e != nil {
		fmt.Println("Config unmarshal failed.")
	}
	return c, true
}

func (c *Config) Save() {
	b, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	err = ioutil.WriteFile(filepath.Join(WORKSPACE_DIR, CONFIG_FILENAME), b, 0777)
	if err != nil {
		log.Fatalf("File write failed: %s", err)
	}
}

// language setting related
// consider changes on multi-languages in the future,
// using api to multipulate the src & tgt settings

func (c *Config) SetSrcLang(src string) {
	c.SrcLang = src
	os.Mkdir(filepath.Join(WORKSPACE_DIR, src), 0777)
	c.Save()
}

func (c *Config) SetTgtLang(tgt string) {
	c.TgtLang = tgt
	os.Mkdir(filepath.Join(WORKSPACE_DIR, tgt), 0777)
	c.Save()
}

func (c *Config) GetSrcLang() string {
	return c.SrcLang
}

func (c *Config) GetTgtLang() string {
	return c.TgtLang
}

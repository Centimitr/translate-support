package support

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	WORKSPACE_DIR   = "translate_workspace"
	CONFIG_FILENAME = "config.json"
)

type Version struct {
	Name string `json:name`
	// GitHash `json:gitHash`
}

type Config struct {
	srcLang  string    `json:source`
	tgtLang  string    `json:target`
	versions []Version `json:versions`
	watch    []string  `json:watch`
}

func Ins() Config {
	var c Config
	_, e := ioutil.ReadFile(filepath.Join(WORKSPACE_DIR, CONFIG_FILENAME))
	if e != nil {
		fmt.Fprintln(os.Stderr, "Config file read err:", e)
	}
	// if e := json.Unmarshal(data, c); e != nil {
	// 	fmt.Fprintln(os.Stderr, "JSON unmarshaling failed: %s", e)
	// }
	return c
}

func (c *Config) Save() {

}

// language setting related
// consider changes on multi-languages in the future,
// using api to multipulate the src & tgt settings

func (c *Config) SetSrcLang(src string) {
	c.srcLang = src
	os.Mkdir(filepath.Join(WORKSPACE_DIR, src), 0777)
	c.Save()
}

func (c *Config) SetTgtLang(tgt string) {
	c.tgtLang = tgt
	os.Mkdir(filepath.Join(WORKSPACE_DIR, tgt), 0777)
	c.Save()
}

func (c *Config) GetSrcLang() string {
	return c.srcLang
}

func (c *Config) GetTgtLang() string {
	return c.tgtLang
}

// version related

func (c *Config) GetVers() []Version {
	return c.versions
}

func (c *Config) GetLatestVer() (Version, error) {
	if len(c.versions) > 0 {
		return c.versions[len(c.versions)-1], nil
	}
	return Version{}, nil
}

func (c *Config) GetPrevVer() (Version, error) {
	if len(c.versions) > 1 {
		return c.versions[len(c.versions)-2], nil
	}
	return Version{}, nil
}

func (c *Config) hasVer(vername string) bool {
	for _, v := range c.versions {
		if v.Name == vername {
			return true
		}
	}
	return false
}

func (c *Config) AddVer(vername string) error {
	if !c.hasVer(vername) {
		c.versions = append(c.versions, Version{vername})
		os.Mkdir(filepath.Join(WORKSPACE_DIR, c.GetSrcLang(), vername), 0777)
		os.Mkdir(filepath.Join(WORKSPACE_DIR, c.GetTgtLang(), vername), 0777)
		c.Save()
		return nil
	}
	// version exist
	return nil
}

func (c *Config) RemoveVer(vername string) error {
	if /*ver exist*/ true {
		for i, v := range c.versions {
			if v.Name == vername {
				copy(c.versions[i:], c.versions[i+1:])
				// check
				c.versions = c.versions[:len(c.versions)-1]
				os.Remove(filepath.Join(WORKSPACE_DIR, c.GetSrcLang(), vername))
				os.Remove(filepath.Join(WORKSPACE_DIR, c.GetTgtLang(), vername))
				c.Save()
				break
			}
		}
		// can't find required vername
	}
	//ver is not exist
	return nil
}

// watch related

func (c *Config) haswatch(filename string) bool {
	for _, fn := range c.watch {
		if fn == filename {
			return true
		}
	}
	return false
}

func (c *Config) GetWatchs() []string {
	return c.watch
}

func (c *Config) AddWatch(filename string) {
	fn := filepath.Clean(filename)
	if c.haswatch(fn) {
		c.watch = append(c.watch, filepath.Clean(fn))
	}
	c.Save()
}

func (c *Config) RemoveWatch(filename string) {
	for i, fn := range c.watch {
		if fn == filename {
			copy(c.watch[i:], c.watch[i+1:])
			// check
			c.watch = c.watch[:len(c.watch)-1]
			break
		}
	}
	c.Save()
}

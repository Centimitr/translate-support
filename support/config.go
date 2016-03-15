package support

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	WORKSPACE_DIR   = "translate_workspace"
	CONFIG_FILENAME = "config.json"
)

type Version struct {
	Name    string   `json:"name"`
	BaseVer string   `json:"basever"`
	Watch   []string `json:"watch"`
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
	_, e := ioutil.ReadFile(filepath.Join(WORKSPACE_DIR, CONFIG_FILENAME))
	// if e != nil {
	// 	fmt.Fprintln(os.Stderr, "Config file read err:", e)
	// }
	// if e := json.Unmarshal(data, c); e != nil {
	// 	fmt.Fprintln(os.Stderr, "JSON unmarshaling failed: %s", e)
	// }
	c.Port = "4567"
	if e != nil {
		return c, false
	} else {
		return c, true
	}
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

// version related

func (c *Config) GetVers() []Version {
	return c.Versions
}

func (c *Config) GetLatestVer() (Version, error) {
	if len(c.Versions) > 0 {
		return c.Versions[len(c.Versions)-1], nil
	}
	return Version{}, nil
}

func (c *Config) GetPrevVer() (Version, error) {
	if len(c.Versions) > 1 {
		return c.Versions[len(c.Versions)-2], nil
	}
	return Version{}, nil
}

func (c *Config) hasVer(vername string) bool {
	for _, v := range c.Versions {
		if v.Name == vername {
			return true
		}
	}
	return false
}

func (c *Config) AddVer(vername string) error {
	if !c.hasVer(vername) {
		c.Versions = append(c.Versions, Version{vername, "", []string{}})
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
		for i, v := range c.Versions {
			if v.Name == vername {
				copy(c.Versions[i:], c.Versions[i+1:])
				// check
				c.Versions = c.Versions[:len(c.Versions)-1]
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

// Watch related

func (c *Config) hasWatch(vername string, filename string) bool {
	for _, v := range c.Versions {
		if vername == v.Name {
			for _, fn := range v.Watch {
				if fn == filename {
					return true
				}
			}
			break
		}
	}
	return false
}

func (c *Config) GetWatchs(vername string) []string {
	for _, v := range c.Versions {
		if vername == v.Name {
			return v.Watch
		}
	}
	return []string{}
}

func (c *Config) AddWatch(vername string, filename string) {
	fn := filepath.Clean(filename)
	if !c.hasWatch(vername, fn) {
		for _, v := range c.Versions {
			if vername == v.Name {
				v.Watch = append(v.Watch, fn)
			}
		}
	}
	c.Save()
}

func (c *Config) RemoveWatch(vername string, filename string) {
	for _, v := range c.Versions {
		if vername == v.Name {
			for i, fn := range v.Watch {
				if fn == filename {
					copy(v.Watch[i:], v.Watch[i+1:])
					// check
					v.Watch = v.Watch[:len(v.Watch)-1]
					break
				}
			}
		}
	}

	c.Save()
}

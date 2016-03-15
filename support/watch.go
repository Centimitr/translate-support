package support

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "os"
	"path/filepath"
	// "strconv"
	// "time"
)

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
		for i, v := range c.Versions {
			if vername == v.Name {
				c.Versions[i].Watch = append(v.Watch, fn)
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
					c.Versions[i].Watch = v.Watch[:len(v.Watch)-1]
					break
				}
			}
		}
	}
	c.Save()
}

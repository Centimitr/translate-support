package support

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "os"
	"github.com/Centimitr/translate-support/tool"
	"path/filepath"
	// "strconv"
	// "time"
)

func (c *Config) hasWatch(vername string, path string) bool {
	for _, v := range c.Versions {
		if vername == v.Name {
			for _, fp := range v.Watch {
				if fp == path {
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

func (c *Config) AddWatch(vername string, path string) {
	fp := filepath.Clean(path)
	if !c.hasWatch(vername, fp) {
		for i, v := range c.Versions {
			if vername == v.Name {
				c.Versions[i].Watch = append(v.Watch, fp)
			}
		}
	}
	c.Save()
}

func (c *Config) RemoveWatch(vername string, path string) {
	for vi, v := range c.Versions {
		if vername == v.Name {
			for i, fp := range v.Watch {
				if fp == path {
					c.Versions[vi].Watch = tool.StringSliceRemove(v.Watch, i)
					break
				}
			}
		}
	}
	c.Save()
}

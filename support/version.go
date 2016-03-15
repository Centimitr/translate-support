package support

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func (c *Config) GetVers() []Version {
	return c.Versions
}

// func (c *Config) GetLatestVer() (Version, error) {
// 	if len(c.Versions) > 0 {
// 		return c.Versions[len(c.Versions)-1], nil
// 	}
// 	return Version{}, nil
// }

// func (c *Config) GetPrevVer() (Version, error) {
// 	if len(c.Versions) > 1 {
// 		return c.Versions[len(c.Versions)-2], nil
// 	}
// 	return Version{}, nil
// }

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
		c.Versions = append(c.Versions, Version{vername, "", []string{}, strconv.Itoa(int(time.Now().UnixNano()))})
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

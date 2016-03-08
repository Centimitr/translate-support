package support

import (
	"os"
	"path/filepath"
)

func Init() (ok bool) {
	e := os.Mkdir(WORKSPACE_DIR, 0777)
	if e != nil {
		return false
	}
	_, e = os.Create(filepath.Join(WORKSPACE_DIR, CONFIG_FILENAME))
	if e != nil {
		return false
	}
	return true
}

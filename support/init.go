package support

import (
	"os"
)

func Init() (ok bool) {
	e := os.Mkdir(WORKSPACE_DIR, 0777)
	if e != nil {
		return false
	}
	// file create is needless because config.Save() will create it if the file is not exist
	// _, e = os.Create(filepath.Join(WORKSPACE_DIR, CONFIG_FILENAME))
	// if e != nil {
	// 	return false
	// }
	var cfg Config
	cfg.Save()
	return true
}

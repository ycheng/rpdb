package dirs

import (
	"os"
	"path/filepath"
)

var (
	// GlobalDataDir the data directory
	GlobalDataDir string
)

func init() {
	dbPath := os.Getenv("RPDB_HOME")
	if len(dbPath) > 0 {
		GlobalDataDir = dbPath
	} else {
		homePath := os.Getenv("HOME")
		GlobalDataDir = filepath.Join(homePath, "/.rpdb")
	}

}

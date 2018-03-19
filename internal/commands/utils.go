package commands

import (
	"os"
	"path/filepath"

	"github.com/cloudflare/cfssl/log"
)

func checkAndCreatePath(path string) {
	//Check if directory exists
	if _, err := os.Stat(filepath.Dir(path)); err != nil {
		log.Fatalf("Cannot create dockumentation, error: %s", err.Error())
	}

	// Create / replace file
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Cannot not create dockumentation, err %s", err.Error())
	}
	defer file.Close()
}

func isValidPath(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

package commands

import "os"

func isValidPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

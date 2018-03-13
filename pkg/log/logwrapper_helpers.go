package log

import (
	"path/filepath"
	"strings"
)

func getFileNameCapitalized(path string) string {
	parent := strings.ToUpper(filepath.Base(filepath.Dir(path)))
	return parent + "|" + strings.ToUpper(strings.Split(filepath.Base(path), ".")[0])
}

package fileutils

import (
	"os"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func AppendSlash(path string) string {
	path = strings.TrimSpace(path)
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return path
}

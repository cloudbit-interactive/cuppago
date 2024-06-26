package cuppago

import (
	"os"
	"path/filepath"
	"strings"
)

func GetRootPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Dir(ex)
	if strings.Contains(strings.ToLower(path), "/temp/") == true || strings.Contains(strings.ToLower(path), "/tmp/") == true || strings.Contains(strings.ToLower(path), "private/") == true || strings.Contains(strings.ToLower(path), "/t/") == true {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return path
	}
	return path
}

func CreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
}

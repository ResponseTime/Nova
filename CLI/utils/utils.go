package utils

import (
	"os"
	"path/filepath"
)

func InternalCreateFolder(path, name string) string {
	err := os.Mkdir(filepath.Join(path, name), 0755)
	if err != nil {
		panic(err)
	}
	return filepath.Join(path, name)
}
func InternalCreateFile(path, name string) string {
	_, err := os.Create(filepath.Join(path, name))
	if err != nil {
		panic(err)
	}
	return filepath.Join(path, name)
}

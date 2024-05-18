package utils

import (
	"os"
	"path/filepath"
)

func InternalCreateFolder(path, name string) error {
	if path == "" {
		return os.Mkdir(name, 0755)
	}
	return os.Mkdir(filepath.Join(path, name), 0755)
}
func InternalCreateFile(path, name string) (*os.File, error) {
	if path == "" {
		return os.Create(name)
	}
	return os.Create(filepath.Join(path, "/", name))
}

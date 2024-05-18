package filehandle

import (
	"path/filepath"

	"github.com/responsetime/Nova/utils"
)

func GENERATE_PACKAGE_JSON(path, project_name string) {
	utils.InternalCreateFile(filepath.Join(path, project_name), "package.json")
}

func CREATE_PROJECT(current_path, project_name string) {
	utils.InternalCreateFolder(current_path, project_name)
}

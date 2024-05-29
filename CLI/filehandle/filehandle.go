package filehandle

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Structure struct {
	Dest     string      `json:"dest"`
	IsFolder bool        `json:"isfolder"`
	Children []Structure `json:"children"`
}

//go:embed filestructre.json
var son []byte
var eve Structure

// files to populate

//go:embed package.json.template
var packagejson []byte

//go:embed App.css
var AppCss []byte

//go:embed App.jsx
var AppJSX []byte

//go:embed index.css
var IndexCss []byte

//go:embed index.html
var IndexHTML []byte

//go:embed main.jsx
var MainJSX []byte

var map_file_to_content_file = map[string][]byte{
	"package.json": packagejson,
	"index.html":   IndexHTML,
	"App.css":      AppCss,
	"App.jsx":      AppJSX,
	"index.css":    IndexCss,
	"main.jsx":     MainJSX,
}

func create_dir_structure(current_path, dest string, eve *Structure, project_name string) {
	if eve.IsFolder {
		os.Mkdir(filepath.Join(current_path, eve.Dest), 0777)
	} else {
		os.Create(filepath.Join(current_path, eve.Dest))
		go func(path string, name string) {
			if name == "package.json" {
				tmpl := template.Must(template.New("json").Parse(string(packagejson)))
				var output bytes.Buffer
				if err := tmpl.Execute(&output, strings.ToLower(project_name)); err != nil {
					log.Fatalf("Error executing template: %v", err)
				}
				if err := os.WriteFile(filepath.Join(path, name), output.Bytes(), 0777); err != nil {
					log.Fatalf("Error writing output file: %v", err)
				}
				return
			}
			os.WriteFile(filepath.Join(path, name), map_file_to_content_file[name], 0777)
		}(current_path, eve.Dest)
	}
	current_path = filepath.Join(current_path, dest)
	for _, itr := range eve.Children {
		create_dir_structure(current_path, itr.Dest, &itr, project_name)
	}
}
func CREATE_PROJECT(current_path, project_name, template, language string) {
	json.Unmarshal(son, &eve)
	eve.Dest = project_name
	create_dir_structure(current_path, project_name, &eve, project_name)

	// path_folder_root := utils.InternalCreateFolder(current_path, project_name)
	// public_folder := utils.InternalCreateFolder(path_folder_root, "public")
	// src_folder := utils.InternalCreateFolder(path_folder_root, "src")
	// package_json := utils.InternalCreateFile(path_folder_root, "package.json")
	// index_html := utils.InternalCreateFile(path_folder_root, "index.html")
	// readme := utils.InternalCreateFile(path_folder_root, "README.md")
	// appCss := utils.InternalCreateFile(src_folder, "App.css")
	// indexCss := utils.InternalCreateFile(src_folder, "index.css")
	// appJSX := utils.InternalCreateFile(src_folder, "App.jsx")
	// mainJSX := utils.InternalCreateFile(src_folder, "main.jsx")
	// fmt.Println(public_folder, package_json, index_html, readme, appCss, indexCss, appJSX, mainJSX)

}

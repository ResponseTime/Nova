package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestInternalCreateFolder(t *testing.T) {
	curDir, _ := os.Getwd()
	var tests = []struct {
		path, folder_name string
		outFile           string
	}{
		{"CLI/Test_folder", "project1", filepath.Join(curDir, "CLI/Test_folder", "project1")},
		{"", "project2", "/project2"},
		{"", "", ""},
		{"CLI/Test_folder", "", ""},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Creating folder %s at path %s ", tt.folder_name, tt.path)
		t.Run(testname, func(t *testing.T) {
			ans := InternalCreateFolder(tt.path, tt.folder_name)
			if ans != tt.outFile {
				t.Errorf("got %v, want %v", ans, tt.outFile)
			}
		})
	}
}

type outFile struct {
	File *os.File
	Err  error
}

func TestInternalCreateFile(t *testing.T) {
	curDir, _ := os.Getwd()
	createTemp := func(path, name string) *outFile {
		file, err := os.CreateTemp(path, name)
		return &outFile{File: file, Err: err}
	}
	var tests = []struct {
		path, file_name string
		outFile         *outFile
	}{
		{"CLI/Test_folder", "testPackageJson.json", createTemp(curDir+"CLI/Test_folder", "testPackageJson.json")},
		{"", "test.txt", createTemp(curDir+"", "testPackageJson.json")},
		{"", "", createTemp(curDir+"", "")},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Creating file %s at path %s ", tt.file_name, tt.path)
		t.Run(testname, func(t *testing.T) {
			file := InternalCreateFile(tt.path, tt.file_name)
			fileToCheck, _ := filepath.Abs(tt.outFile.File.Name())
			if file != fileToCheck {
				t.Errorf("got %v, want %v", file, tt.outFile.File.Name())
			}
		})
	}
}

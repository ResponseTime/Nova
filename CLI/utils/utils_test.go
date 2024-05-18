package utils

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestInternalCreateFolder(t *testing.T) {
	var tests = []struct {
		path, folder_name string
		err               error
	}{
		{"CLI/Test_folder", "project1", nil},
		{"", "project2", nil},
		{"", "", errors.New("Enter Folder Name")},
		{"CLI/Test_folder", "", errors.New("Enter Folder Name")},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Creating folder %s at path %s ", tt.folder_name, tt.path)
		t.Run(testname, func(t *testing.T) {
			ans := InternalCreateFolder(tt.path, tt.folder_name)
			if ans != tt.err {
				t.Errorf("got %v, want %v", ans, tt.err)
			}
		})
	}
}

type outFile struct {
	File *os.File
	Err  error
}

func TestInternalCreateFile(t *testing.T) {
	createTemp := func(path, name string) *outFile {
		file, err := os.CreateTemp(path, name)
		return &outFile{File: file, Err: err}
	}
	var tests = []struct {
		path, file_name string
		outFile         *outFile
	}{
		{"CLI/Test_folder", "testPackageJson.json", createTemp("CLI/Test_folder", "testPackageJson.json")},
		{"", "test.txt", createTemp("", "testPackageJson.json")},
		{"", "", createTemp("", "")},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Creating file %s at path %s ", tt.file_name, tt.path)
		t.Run(testname, func(t *testing.T) {
			file, err := InternalCreateFile(tt.path, tt.file_name)
			if err != nil {
				if err == tt.outFile.Err {
					t.Errorf("got %v, want %v", err, tt.outFile.Err)
				}
			} else {
				if file != tt.outFile.File {
					t.Errorf("got %v, want %v", file, tt.outFile.File)
				}
			}
		})
	}
}

package utils

import (
	"os"
	"path/filepath"
)

//IFile file interface
type IFile interface {
	Path() string
}

//File utils for files (configs)
type File struct {
	FilePath string
}

//Path returns absolute file path
func (file File) Path() string {
	dir := filepath.Dir(os.Args[0])
	absPath, _ := filepath.Abs(dir + file.FilePath)
	return absPath
}

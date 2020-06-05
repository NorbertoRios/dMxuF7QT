package utils

import "path/filepath"

type IFile interface{
	Path() string
}

//FileUtils utils for files (configs)
type File struct {
	Filename string
}

//Path returns absolute file path
func (file File) Path() string {
	absPath, _ := filepath.Abs(file.Filename)
	return absPath
}

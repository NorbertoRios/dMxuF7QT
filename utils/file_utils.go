package utils

import "path/filepath"

//FileUtils utils for files (configs)
type FileUtils struct {
	Filename string
}

//Path returns absolute file path
func (utils *FileUtils) Path() string {
	absPath, _ := filepath.Abs(utils.Filename)
	return absPath
}

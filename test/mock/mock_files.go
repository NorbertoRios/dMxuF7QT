package mock

import "path/filepath"

//File mock file for tests
type File struct {
	FilePath string
}

//Path returns absolute file path
func (file File) Path() string {
	absPath, _ := filepath.Abs(file.FilePath)
	return absPath
}

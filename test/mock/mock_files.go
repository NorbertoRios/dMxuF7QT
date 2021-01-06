package mock

import "path/filepath"

//File mock file for tests
type File struct {
	fileDest string
	dir      string
}

//Path returns absolute file path
func (file File) Path() string {
	absPath, _ := filepath.Abs(filepath.Dir(file.dir) + file.fileDest)
	return absPath
}

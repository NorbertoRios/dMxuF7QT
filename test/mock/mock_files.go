package mock

import "path/filepath"

//NewFile ...
func NewFile(_dir, filename string) *File {
	return &File{
		dir:      _dir,
		fileName: filename,
	}
}

//File mock file for tests
type File struct {
	fileName string
	dir      string
}

//Path returns absolute file path
func (file File) Path() string {
	absPath, _ := filepath.Abs(filepath.Dir(file.dir) + file.fileName)
	return absPath
}

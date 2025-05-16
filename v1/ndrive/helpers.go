package ndrive

import (
	"os"
	"path"
)

func CreateFileAt(filename string, workingdir string, relativePath string, data []byte) error {

	fullPath := path.Join(workingdir, relativePath, filename)

	_, err := os.Create(fullPath)
	if err != nil {
		return ErrorCreateFileFailed
	}
	return nil
}

func Fullpath(workingdir string, relativePath string) string {
	return path.Join(workingdir, relativePath)
}
func FileExists(fullPath, filename string) bool {
	_, err := os.Stat(path.Join(fullPath, filename))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func MakeDir(workingdir string, relativePath string) error {
	fullPath := path.Join(workingdir, relativePath)
	err := os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

package fun

import "path"









func Fullpath(workingdir string, relativePath string) string {
	return path.Join(workingdir, relativePath)
}

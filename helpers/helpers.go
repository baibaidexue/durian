package helpers

import (
	"os"
	"path/filepath"
)

// EnsurePath ensures the given filesystem path exists, if not it will create it.
func EnsurePath(pathName string) error {
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		err = os.MkdirAll(pathName, 0755)
		if err != nil {

			return err
		}
	}
	return nil
}

// FileExists checks whether a file or folder exists in the filesystem, will follow symlinks and ensures the target exists too.
func FileExists(pathName string) bool {
	fi, err := os.Lstat(pathName)
	if err != nil {
		return false
	}
	// Is a symlink
	if fi.Mode()&os.ModeSymlink != 0 {

		p, err := filepath.EvalSymlinks(pathName)
		if err == nil {
			return FileExists(p)
		}
		return false
	}
	return true
}
package util

import (
	"io"
	"os"
)

func OpenFile(filename string, appendFlag bool) (*os.File, error) {
	flag := os.O_RDWR | os.O_CREATE
	if appendFlag {
		flag |= os.O_APPEND
	}

	return os.OpenFile(JoinPath(ROOT_DIR, filename), flag, os.ModePerm)
}

func ReadFile(filename string) ([]byte, error) {
	file, err := OpenFile(filename, false)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}

func ReadDir(path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}

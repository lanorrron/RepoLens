package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func VerifyPath(path string) (string, error) {

	if strings.TrimSpace(path) == "" {
		return "", errors.New("empty path")

	}
	cleanedPath := filepath.Clean(path)

	realPath, err := filepath.EvalSymlinks(cleanedPath)
	if err != nil {
		return "", errors.New("invalid route or not secure")
	}
	info, err := os.Stat(realPath)
	if err != nil {
		return "", errors.New("path does not exit")
	}

	if !info.IsDir() {
		return "", errors.New("path is not a directory")
	}

	return realPath, nil
}

package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	MaxFileSize = 1024 * 1024 * 10 // 10MB limit,
)

// reads the file securely
func ReadFileSecurely(path string) (string, error) {
	cleanPath := filepath.Clean(path)
	if strings.Contains(cleanPath, "..") { // checks if the user is trying to access to files out of the current directory
		return "", fmt.Errorf("path traversal not allowed")
	}

	if !strings.HasSuffix(cleanPath, ".txt") { // checks if the file suffix is txt
		return "", fmt.Errorf("file is not of type txt")
	}

	fileInfo, err := os.Stat(cleanPath)
	if err != nil { // checks if the stat file exists
		return "", fmt.Errorf("file was not found: %v", err)
	}

	if !fileInfo.Mode().IsRegular() { // checks if the file represents a regular file (not in directories, symbolic links, devices)
		return "", fmt.Errorf("not a regular file")
	}

	if fileInfo.Size() > MaxFileSize { // checks if file size does not exceed the permitted size
		return "", fmt.Errorf("file is too large (max %d) bytes", MaxFileSize)
	}

	content, err := os.ReadFile(cleanPath)
	if err != nil { // reads the file
		return "", fmt.Errorf("file could not be read: %v", err)
	}

	return string(content), nil
}

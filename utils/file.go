package utils

import (
	"fmt"
	"io"
	"os"
)

// CreateTempDir creates a temporary directory for code execution.
func CreateTempDir() (string, func(), error) {
	dir, err := os.MkdirTemp("", "execution-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating temporary directory: %w", err)
	}
	return dir, func() { os.RemoveAll(dir) }, nil
}

// WriteCodeToFile writes the provided code to the specified file path.
func WriteCodeToFile(filePath, code string) error {
	if err := os.WriteFile(filePath, []byte(code), 0644); err != nil {
		return fmt.Errorf("error writing to file %s: %w", filePath, err)
	}
	return nil
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	dstInfo, err := os.Stat(dst)
	if err == nil && os.SameFile(srcInfo, dstInfo) {
		return nil
	}
	if err := os.Link(src, dst); err == nil {
		return nil
	}
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return dstFile.Close()
}

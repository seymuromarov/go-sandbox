// dockerutils/executor.go
package dockerutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// getDockerCommand returns the Docker command to execute based on the file extension.
func GetDockerCommand(extension, filePath string) string {
	baseName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	workspaceDir := GetSetting("WORKSPACE_DIR")

	switch extension {
	case ".py":
		return fmt.Sprintf("python3 %s/%s", workspaceDir, filepath.Base(filePath))
	case ".js":
		return fmt.Sprintf("node %s/%s", workspaceDir, filepath.Base(filePath))
	case ".ts":
		return fmt.Sprintf("tsc %s/%s.ts && node %s/%s.js", workspaceDir, baseName, workspaceDir, baseName)
	case ".java":
		return fmt.Sprintf("javac %s/%s.java && java -cp %s %s", workspaceDir, baseName, workspaceDir, baseName)
	case ".cpp":
		return fmt.Sprintf("g++ %s/%s.cpp -o %s/%s && %s/%s", workspaceDir, baseName, workspaceDir, baseName, workspaceDir, baseName)
	case ".go":
		return fmt.Sprintf("go run %s/%s.go", workspaceDir, baseName)
	default:
		return ""
	}
}

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
}

// GetSetting retrieves the value of a setting from the environment variables.
func GetSetting(key string) string {
	return os.Getenv(key)
}

package main

import (
	"fmt"
	"main/dockerutils"
	"os"
	"path/filepath"
)

func main() {
	dockerutils.LoadEnv()
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <language extension> <file path>\nExample: go run . .go ./examples/fibonacci.go")
		os.Exit(1)
	}

	langExt := os.Args[1]
	filePath := os.Args[2]

	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Printf("Error getting file path: %v\n", err)
		os.Exit(1)
	}

	dockerutils.ExecuteCode(langExt, absFilePath)
}

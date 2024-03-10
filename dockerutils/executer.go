package dockerutils

import (
	"bytes"
	"fmt"
	"main/utils"
	"os"
	"os/exec"
	"path/filepath"
)

// buildDockerImage builds the Docker image used for code execution.
func buildDockerImage() error {
	dockerImageName := fmt.Sprintf("%s:%s", GetSetting("DOCKER_IMAGE_NAME"), GetSetting("DOCKER_IMAGE_VERSION"))
	cmd := exec.Command("docker", "build", "-t", dockerImageName, ".")
	printDockerBuildConsole := GetSetting("PRINT_DOCKER_BUILD_CONSOLE")

	if printDockerBuildConsole == "true" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

// runCodeInDocker executes the code in a Docker container based on the file extension.
func runCodeInDocker(extension, dir, filename string) error {
	workspaceDir := GetSetting("WORKSPACE_DIR")
	command := GetDockerCommand(extension, filename)
	if command == "" {
		return fmt.Errorf("unsupported file extension: %s", extension)
	}

	time_commandPath := filepath.Join(dir, "time_command.go")
	if err := utils.CopyFile(filepath.Join("helpers", "time_command.go"), time_commandPath); err != nil {
		return fmt.Errorf("failed to copy time_command.go to tmp dir: %v", err)
	}

	printExecutionTime := GetSetting("PRINT_EXECUTION_TIME")

	if printExecutionTime == "true" {

		command = fmt.Sprintf("go run %s/time_command.go %s", workspaceDir, command)
	}
	wrappedCommand := fmt.Sprintf("sh -c 'cd %s && %s'", workspaceDir, command)

	dockerImageName := fmt.Sprintf("%s:%s", GetSetting("DOCKER_IMAGE_NAME"), GetSetting("DOCKER_IMAGE_VERSION"))

	timeLimit := GetSetting("DEFAULT_TIME_LIMIT") // Assuming GetSetting can retrieve environment variables
	memoryLimit := GetSetting("DEFAULT_MEMORY_LIMIT")

	// Mount the directory to workspaceDir in the container.
	runCmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:%s", dir, workspaceDir),
		"--memory", memoryLimit,
		dockerImageName, "sh", "-c", fmt.Sprintf("timeout %s %s", timeLimit, wrappedCommand))

	var runOut bytes.Buffer
	runCmd.Stdout = &runOut
	runCmd.Stderr = &runOut

	if err := runCmd.Run(); err != nil {
		fmt.Printf("Container output: %s\n", runOut.String())
		return err
	}
	fmt.Printf("Execution output: \n%s", runOut.String())

	// Clean up: Remove the temporary time_command.go after execution
	if err := os.Remove(time_commandPath); err != nil {
		fmt.Printf("Warning: Failed to remove temporary file %s: %v\n", time_commandPath, err)
	}

	return nil
}

// ExecuteCode runs the provided code file in a Docker container based on the file extension.
func ExecuteCode(fileExtension, filePath string) {
	dir, _ := filepath.Split(filePath)
	filename := filepath.Base(filePath)

	if err := buildDockerImage(); err != nil {
		fmt.Printf("Error building Docker image: %v\n", err)
		return
	}
	if err := runCodeInDocker(fileExtension, dir, filename); err != nil {
		fmt.Printf("Error running code in Docker: %v\n", err)
		return
	}
}

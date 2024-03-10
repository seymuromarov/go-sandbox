package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	startTime := time.Now()
	// Execute the command passed as arguments to the script
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	elapsedTime := time.Since(startTime)

	fmt.Printf("Execution Time: %.2f ms\n", elapsedTime.Seconds()*1000)
}

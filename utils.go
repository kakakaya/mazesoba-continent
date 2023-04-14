package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func openDirectory(path string) error {
	var command string
	switch runtime.GOOS {
	case "darwin":
		command = "open"
	case "linux":
		command = "xdg-open"
	case "windows":
		command = "explorer.exe"
	default:
		return fmt.Errorf("unsupported platform")
	}
	cmd := exec.Command(command, path)
	return cmd.Run()
}

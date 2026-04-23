package executor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var blockedPatterns = []string{
	"os.Remove",
	"os.RemoveAll",
	"exec.Command",
	"syscall",
	"unsafe",
	"net/http",
}

func isSafe(code string) bool {
	for _, b := range blockedPatterns {
		if strings.Contains(code, b) {
			return false
		}
	}

	if !strings.Contains(code, "package main") ||
		!strings.Contains(code, "func main") {
		return false
	}

	return true
}

func Run(code string) (string, error) {
	if !isSafe(code) {
		return "", fmt.Errorf("blocked unsafe Go code")
	}

	file, err := os.CreateTemp("", "agent-*.go")
	if err != nil {
		return "", fmt.Errorf("temp file error: %w", err)
	}
	defer os.Remove(file.Name())

	if err := os.WriteFile(file.Name(), []byte(code), 0644); err != nil {
		return "", fmt.Errorf("write file error: %w", err)
	}

	if _, err := exec.LookPath("goimports"); err == nil {
		cmdFmt := exec.Command("goimports", "-w", file.Name())
		if err := cmdFmt.Run(); err != nil {
			fmt.Println("warning: goimports failed:", err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", file.Name())
	out, err := cmd.CombinedOutput()

	return string(out), err
}

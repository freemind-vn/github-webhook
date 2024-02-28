package helper

import (
	"log/slog"
	"os"
	"os/exec"
	"time"
)

var debounced func(f func())

func InitCommand(workingDir string, timeout int64) {
	debounced = Debouncer(time.Duration(timeout * int64(time.Second)))
}

// Runs multiple commands in a single shell instance
func RunCommand(dir, commands string) {
	slog.Info("RunCommand", "cmd", commands)
	cmd := exec.Command("/bin/sh", "-c", commands)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		slog.Error("RunCommand", "err", err)
	}
}

// Commits the changes
func RunDebouncedCommand(dir, commands string) {
	slog.Info("RunDebouncedCommand", "cmd", commands)
	debounced(func() {
		RunCommand(dir, commands)
	})
}

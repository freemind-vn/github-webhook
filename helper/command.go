package helper

import (
	"os"
	"os/exec"
	"time"

	"github.com/rs/zerolog/log"
)

var debounced func(f func())

func InitCommand(workingDir string, timeout int64) {
	debounced = Debouncer(time.Duration(timeout * int64(time.Second)))
}

// Runs multiple commands in a single shell instance
func RunCommand(dir, commands string) {
	log.Info().Msgf("shell commands: %s", commands)
	cmd := exec.Command("/bin/sh", "-c", commands)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Error().Msgf("run command: %s", err)
	}
}

// Commits the changes
func RunDebouncedCommand(dir, commands string) {
	log.Info().Msgf("shell (debounced): %s", commands)
	debounced(func() {
		RunCommand(dir, commands)
	})
}

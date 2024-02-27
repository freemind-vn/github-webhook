package cmd

import (
	"fmt"
	"os"
	"runtime"

	"freemind.com/webhook/service"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "webhook",
	Short:   "GitHub Webhook Event Notifications",
	Version: fmt.Sprintf("v%s-%s/%s %s/%s BuildDate=%s\n", service.Version, service.Hash, service.Branch, runtime.GOOS, runtime.GOARCH, service.BuildDate),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Return the root command
func GetRootCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

package cmd

import (
	"github.com/spf13/cobra"

	"freemind.com/webhook/service"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return service.ServeHTTP(config)
		},
	}
	config = "data/configs/config.yaml"
)

// Return the serve command
func GetServeCmd() *cobra.Command {
	return serveCmd
}

func init() {
	serveCmd.Flags().StringVarP(&config, "config", "c", config, "path to the config file")
	rootCmd.AddCommand(serveCmd)
}

package cmd

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"freemind.com/webhook/service"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve()
	},
}

// Return the serve command
func GetServeCmd() *cobra.Command {
	return serveCmd
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// Start the server
func serve() error {
	log.Info().Msgf("serve on: %v", service.ServerPort)

	s := &service.Service{}
	return http.ListenAndServe(service.ServerPort, s)
}

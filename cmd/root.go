package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "designpilot",
	Short: "A brief description of your application",
	Long: `Design Pilot is an AI-powered branding assistant.
This is the CLI frontend that allows users to interact with certain parts of the service.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}

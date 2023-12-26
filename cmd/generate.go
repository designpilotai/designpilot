/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"designpilot.ai/designpilot-cli/internal/env"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate assets using Design Pilot.",
	Long: `Generate assets using Design Pilot.

Example:
designpilot generate --logo`,
	Run: func(cmd *cobra.Command, args []string) {
		key := env.GetApiKey()
		fmt.Println("Got API Key: ", key)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// generateCmd.Flags().BoolP("logo", "l", false, "Generated asset will be a logo")
	generateCmd.Flags().StringP("prompt", "p", "", "Prompt for asset")
}

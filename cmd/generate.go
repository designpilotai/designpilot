/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"designpilot.ai/designpilot/internal/env"
	"designpilot.ai/designpilot/internal/network"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate assets using Design Pilot.",
	Long: `Generate assets using Design Pilot.

Example:
designpilot generate -p="A pig in a teacup"`,
	Run: func(cmd *cobra.Command, args []string) {
		key := env.GetApiKey()
		origin := env.GetOrigin()

		reqUrl := origin + "/api/cli/generate"

		prompt := cmd.Flag("prompt").Value.String()
		if prompt == "" {
			log.Fatalln("Must set a prompt")
		}

		logo := cmd.Flag("logo").Value.String()

		data := url.Values{
			"prompt": {prompt},
			"logo":   {logo},
		}

		buffer := bytes.NewBufferString(data.Encode())

		req, err := http.NewRequest("POST", reqUrl, buffer)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "Bearer "+key)

		httpClient := network.HttpClient()

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalln(resp.Status)
		}

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(origin + string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolP("logo", "l", false, "Generated asset will be a logo")
	generateCmd.Flags().StringP("prompt", "p", "", "Prompt for asset")
}

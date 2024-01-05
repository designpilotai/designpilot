package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/designpilotai/designpilot/internal/env"
	"github.com/designpilotai/designpilot/internal/network"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate assets using Design Pilot.",
	Long: `Generate assets using Design Pilot.

Example:
designpilot generate -p="A pig in a teacup"`,
	Run: func(cmd *cobra.Command, args []string) {
		quiet := cmd.Flag("quiet").Value.String() == "true"

		if !quiet {
			fmt.Println("Reading API key...")
		}

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

		if !quiet {
			fmt.Println("Making request to server...")
		}

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

		if !quiet {
			fmt.Println("All done! Get you image here ↓ Tip: replace 'fullres' with '256' for a low-res version.")
		}

		fmt.Println(origin + string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolP("logo", "l", false, "Generated asset will be a logo")
	generateCmd.Flags().StringP("prompt", "p", "", "Prompt for asset")
	generateCmd.Flags().BoolP("quiet", "q", false, "Disables status logs")
}

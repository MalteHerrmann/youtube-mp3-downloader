/*
Copyright © 2024 Malte Herrmann (malteherrmann.mail@gmail.com)
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	binaryName = "ytdl"
)

var (
	outputDir string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   fmt.Sprintf("%s URL", binaryName),
	Args:  cobra.RangeArgs(0, 1),
	Short: "A tool to download audio files from YouTube",
	Long: `This tool is used to download audio files from YouTube.
  It will download a given URL and provide a way for users to set the desired title and output directory.
  It also includes the option of automatically assigning ID3 tags for the exported audio files.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		var link string

		if len(args) == 0 {
			var err error
			link, err = getVideoLink()
			if err != nil {
				return fmt.Errorf("failed to get video link: %w", err)
			}
		} else {
			link = args[0]
		}

		return Entrypoint(link, outputDir)
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	userHome, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	defaultOutputDir := filepath.Join(userHome, "Music", "YouTube-Downloads")
	rootCmd.PersistentFlags().StringVar(&outputDir, "output-dir", defaultOutputDir, "select the base output directory")
}

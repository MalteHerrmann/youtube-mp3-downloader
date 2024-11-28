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
  Args: cobra.ExactArgs(1),
	Short: "A tool to download audio files from YouTube",
	Long: `This tool is used to download audio files from YouTube.
  It will download a given URL and provide a way for users to set the desired title and output directory.
  It also includes the option of automatically assigning ID3 tags for the exported audio files.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
    if len(args) != 1 {
      return fmt.Errorf("expected 1 argument; got: %d", len(args))
    }

    return Entrypoint(args[0], outputDir)
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



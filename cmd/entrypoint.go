package cmd

import (
	"fmt"
	"path/filepath"

	envinfo "github.com/MalteHerrmann/yt-downloader/internal/env"
	"github.com/MalteHerrmann/yt-downloader/internal/fs"
	"github.com/MalteHerrmann/yt-downloader/internal/youtube"
)

// Entrypoint is the main execution function for the bare command.
func Entrypoint(url, outputDir string) error {
  if err := envinfo.CheckEnvironment(); err != nil {
    return fmt.Errorf("failed to validate environment: %w", err)
  }
  
  if err := fs.CheckDirExists(outputDir); err != nil {
    return err
  }

  urlToDownload, err := youtube.ValidateURL(url)
  if err != nil {
    return fmt.Errorf("url failed validation: %w", err)
  }

  monthDir := fs.GetCurrentMonthDir()
  targetDir := filepath.Join(outputDir, monthDir)

  println("downloading ", urlToDownload.String())
  println(" -> into ", targetDir)

  return nil
}


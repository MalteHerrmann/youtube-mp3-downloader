package env

import (
	"fmt"
	"os/exec"

	"github.com/MalteHerrmann/yt-downloader/internal/youtube"
)

// CheckEnvironment makes sure that the required dependencies are set up
// as expected in the execution context.
func CheckEnvironment() error {
  if err := exec.Command(youtube.DownloaderBinary, "--version").Run(); err != nil {
    return fmt.Errorf("%s binary not found: %w", youtube.DownloaderBinary, err)
  }

  return nil
}

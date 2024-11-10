package env

import (
  "fmt"
  "os/exec"
)

// CheckEnvironment makes sure that the required dependencies are set up
// as expected in the execution context.
func CheckEnvironment() error {
  if err := exec.Command("yt-dlp", "--version").Run(); err != nil {
    return fmt.Errorf("yt-dlp binary not found: %w", err)
  }

  return nil
}

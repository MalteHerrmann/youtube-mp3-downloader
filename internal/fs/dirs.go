package fs

import (
  "fmt"
  "os"
  "time"
)

func CheckDirExists(dir string) error {
  fi, err := os.Stat(dir)
  if  err != nil {
    return fmt.Errorf("output directory not found: %w", err)
  }

  if !fi.IsDir() {
    return fmt.Errorf("configured output dir exists but is not a directory: %q", dir)
  }

  return nil
}

func GetCurrentMonthDir() string {
  now := time.Now()
  return fmt.Sprintf("%04d_%02d", now.Year(), now.Month())
}

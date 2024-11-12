package youtube

import (
	"fmt"
  "path/filepath"
  "os/exec"
)

const DownloaderBinary = "yt-dlp"

func DownloadWithYTDLP(videoData *VideoData, exportDir string) error {
  exportFile := filepath.Join(exportDir, getFilename(videoData.Info))

	err := exec.Command(
		DownloaderBinary,
		"-x",
		"--audio-format",
		"mp3",
		videoData.URL.ParsedURL.String(),
		"--output",
		exportFile,
	).Run()
	if err != nil {
		return fmt.Errorf("failed to download with %s: %w", DownloaderBinary, err)
	}

	return nil
}

func getFilename(vi *VideoInfo) string {
  // TODO: check if the video title already contains the expected artist format
  return fmt.Sprintf("%s - %s.mp3", vi.Author, vi.Title)
}

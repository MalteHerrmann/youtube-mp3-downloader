package youtube

import (
	"fmt"
  "os/exec"
)

const DownloaderBinary = "yt-dlp"

func DownloadWithYTDLP(videoData *VideoData, exportFile string) error {
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


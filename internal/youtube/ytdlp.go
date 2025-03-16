package youtube

import (
	"fmt"
	"os/exec"
)

const DownloaderBinary = "yt-dlp"

func DownloadWithYTDLP(videoData *VideoData, exportFile string) error {
	_, err := exec.Command(
		DownloaderBinary,
		"-x",
		"--audio-format",
		"mp3",
		videoData.URL.ParsedURL.String(),
		"--output",
		exportFile,
	).Output()
	if err != nil {
		exErr, ok := err.(*exec.ExitError)
		if ok {
			// NOTE: since exErr.Error() only returns the status code instead of the contents of Stderr,
			// we are returning exErr.StdErr here
			return fmt.Errorf("failed to download with %s: %s", DownloaderBinary, string(exErr.Stderr))
		}

		return fmt.Errorf("failed to download with %s: %w", DownloaderBinary, err)

	}

	return nil
}

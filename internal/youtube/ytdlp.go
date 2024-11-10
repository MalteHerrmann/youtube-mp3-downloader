package youtube

import (
	"fmt"
	"net/url"
	"os/exec"
)

const DownloaderBinary = "yt-dlp"

func DownloadWithYTDLP(videoURL *url.URL, exportDir string) error {
	// TODO: read video metadata and derive title, and from that track and artist, potentially also label with regex [0-9A-Z]+[0-9]+
	// TODO: use youtube Golang package e.g. to sanitize the filename and potentially download too: https://github.com/kkdai/youtube/blob/master/downloader/file_utils.go#L51
	err := exec.Command(
		DownloaderBinary,
		"-x",
		"--audio-format",
		"mp3",
		videoURL.String(),
		"--output",
		// TODO: this should not only be the dir but rather the full file name derived from the title
		exportDir,
	).Run()
	if err != nil {
		return fmt.Errorf("failed to download with %s: %w", DownloaderBinary, err)
	}

	return nil
}

package cmd

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"

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

	videoData, err := youtube.GetVideoData(url)
	if err != nil {
		return fmt.Errorf("failed to get video data: %w", err)
	}

	println("got video title: ", videoData.Info.Title)
	println("got video uploader: ", videoData.Info.Author)

	monthDir := fs.GetCurrentMonthDir()
	targetDir := filepath.Join(outputDir, monthDir)
	fileName, err := getFilename(videoData.Info)
	if err != nil {
		return fmt.Errorf("failed to get filename: %w", err)
	}

	exportFile := filepath.Join(targetDir, fileName)

	println("downloading ", videoData.URL.ParsedURL.String())
	println(" -> into ", exportFile)

	return youtube.DownloadWithYTDLP(videoData, exportFile)
}

// getVideoLink asks the user for the video link
func getVideoLink() (string, error) {
	videoLinkPrompt := promptui.Prompt{
		Label: "Video Link",
	}

	return videoLinkPrompt.Run()
}

// getFilename cleans the video info and prompts the user for the
// artist and title to be used in the filename.
func getFilename(vi *youtube.VideoInfo) (string, error) {
	artist := regexp.MustCompile(`-\s*[tT]opic`).ReplaceAllString(vi.Author, "")
	title := vi.Title

	titleMatch := regexp.MustCompile(`([0-9a-zA-Z\s]+)\s*-\s*([0-9a-zA-Z\s]+)`).FindStringSubmatch(vi.Title)
	if len(titleMatch) == 3 {
		artist = titleMatch[1]
		title = titleMatch[2]
	}

	artistPrompt := promptui.Prompt{
		Label:   "Artist",
		Default: strings.TrimSpace(artist),
	}

	var err error
	artist, err = artistPrompt.Run()
	if err != nil {
		return "", err
	}

	titlePrompt := promptui.Prompt{
		Label:   "Title",
		Default: title,
	}

	title, err = titlePrompt.Run()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s - %s.mp3", artist, title), nil
}

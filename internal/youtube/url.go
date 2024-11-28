package youtube

import (
	"fmt"
	"net/url"
	"regexp"
)

type youtubeURL struct {
	ParsedURL *url.URL
	VideoID   string
}

// parseYouTubeURL parses the given raw YouTube url and returns a
// youtubeURL struct, which contains a cleaned URL and the video ID.
func parseYouTubeURL(rawURL string) (youtubeURL, error) {
	videoPattern := regexp.MustCompile(`(https://www\.youtube\.com/watch\?v=)([a-zA-Z0-9\-_]+)(&list.*){0,1}$`)
	submatches := videoPattern.FindStringSubmatch(rawURL)
	if len(submatches) < 3 {
		return youtubeURL{}, fmt.Errorf("expected youtube watch link; got: %s", rawURL)
	}

	videoID := submatches[2]
	cleanURL, err := url.Parse(submatches[1] + videoID)
	if err != nil {
		return youtubeURL{}, err
	}

	return youtubeURL{
		ParsedURL: cleanURL,
		VideoID:   videoID,
	}, nil
}

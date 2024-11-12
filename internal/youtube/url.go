package youtube

import (
  "fmt"
  "net/url"
  "regexp"
)

type youtubeURL struct {
  ParsedURL *url.URL
  VideoID string
}

func parseYouTubeURL(rawURL string) (youtubeURL, error) {
  parsedURL, err := url.Parse(rawURL)
  if err != nil {
    return youtubeURL{}, err
  }
  
  videoID, err := getVideoIDFromURL(parsedURL)
  if err != nil {
    return youtubeURL{}, err
  }

  return youtubeURL{
    ParsedURL: parsedURL,
    VideoID: videoID,
  }, nil
}

// getVideoIDFromURL returns the stripped video ID from the given
// YouTube url.
//
// NOTE: the URL was already checked to be a valid video link.
func getVideoIDFromURL(videoURL *url.URL) (string, error) {
  videoPattern := regexp.MustCompile(`https://www\.youtube\.com/watch\?v=([a-zA-Z0-9]+)`)
  subMatches := videoPattern.FindStringSubmatch(videoURL.String())

  if len(subMatches) != 2 {
    return "", fmt.Errorf("failed to match video link; expected youtube watch link; got: %s", videoURL.String())
  }

  return subMatches[1], nil
}


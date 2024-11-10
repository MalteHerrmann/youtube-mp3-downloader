package youtube

import (
  "fmt"
  "net/url"
  "strings"
)

func ValidateURL(rawURL string) (*url.URL, error) {
  parsedURL, err := url.Parse(rawURL)
  if err != nil {
    return nil, err
  }
  
  if !strings.Contains(parsedURL.Hostname(), "youtube.com") {
    return nil, fmt.Errorf("expected youtube link; got: %s", rawURL)
  }

  return parsedURL, nil
}


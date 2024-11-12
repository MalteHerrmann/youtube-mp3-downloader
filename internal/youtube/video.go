package youtube

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type VideoInfo struct {
  Title string `json:"title"`
  Author string `json:"author_name"`
}

func getVideoInfo(videoURL youtubeURL) (*VideoInfo, error) {
  // NOTE: we're using the oembed API here which should offer free access to title and uploader of the video
  url := fmt.Sprintf("https://www.youtube.com/oembed?url=%s&format=json", videoURL.ParsedURL)

  resp, err := http.Get(url)
  if err != nil {
      return nil, err
  }
  defer resp.Body.Close()

  var data VideoInfo
  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
      return nil, err
  }

  return &data, nil
}

type VideoData struct {
  Info *VideoInfo
  URL youtubeURL
}

func GetVideoData(rawURL string) (*VideoData, error) {
  urlToDownload, err := parseYouTubeURL(rawURL)
  if err != nil {
    return nil, err
  }

  info, err := getVideoInfo(urlToDownload)
  if err != nil {
    return nil, err
  }

  return &VideoData{
    Info: info,
    URL: urlToDownload,
  }, nil
}

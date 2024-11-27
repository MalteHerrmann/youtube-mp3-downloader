package youtube

import (
	"strings"
	"testing"
)

func TestParseYoutubeURL(t *testing.T) {
	testcases := []struct {
		name       string
		url        string
		expError   string
		expVideoID string
	}{
		{name: "pass - valid youtube video link", url: "https://www.youtube.com/watch?v=Mo3pqAuGnDo", expVideoID: "Mo3pqAuGnDo"},
		{name: "pass - valid youtube video link with list", url: "https://www.youtube.com/watch?v=Mo3pqAuGnDo&list=PL1234567890", expVideoID: "Mo3pqAuGnDo"},
		{name: "fail - not a youtube link", url: "https://google.com", expError: "expected youtube watch link; got: https://google.com"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			parsedURL, err := parseYouTubeURL(tc.url)

			switch {
			case tc.expError == "" && err != nil:
				t.Fatalf("expected no error; got: %s", err)
			case tc.expError == "" && err == nil && tc.expVideoID == parsedURL.VideoID:
				// success case
			case tc.expError == "" && err == nil && tc.expVideoID != parsedURL.VideoID:
				t.Fatalf("expected different video id; expected: %s; got: %s", tc.expVideoID, parsedURL.VideoID)
			case tc.expError != "" && err == nil:
				t.Fatal("expected error; got: nil")
			case tc.expError != "" && !strings.Contains(err.Error(), tc.expError):
				t.Fatalf("expected different error; expected: %q; got: %q", tc.expError, err.Error())
			case tc.expError != "" && strings.Contains(err.Error(), tc.expError):
				// success case
			default:
				t.Fatalf("unexpected combination; err: %s", err)
			}
		})
	}
}

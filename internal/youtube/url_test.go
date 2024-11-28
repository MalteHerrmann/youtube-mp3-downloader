package youtube

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

			if tc.expError == "" {
				assert.NoError(t, err, "expected no error; got: %s", err)
				assert.Equal(t, tc.expVideoID, parsedURL.VideoID, "expected video id: %s; got: %s", tc.expVideoID, parsedURL.VideoID)
			} else {
				assert.Error(t, err, "expected error; got: nil")
				assert.Contains(t, err.Error(), tc.expError, "expected error: %s; got: %s", tc.expError, err)
			}
		})
	}
}

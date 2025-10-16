package pkg

import (
	"io"
	"net/http"
	"strings"
	"time"
)

func ShortDuration(d time.Duration) string {
	s := d.String()

	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}

	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}

	return s
}

func SendRequest(req *http.Request) ([]byte, error) {
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	userData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

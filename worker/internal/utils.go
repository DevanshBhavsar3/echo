package internal

import (
	"net/http"
	"time"

	"github.com/DevanshBhavsar3/echo/common/db/store"
)

func Ping(url string) (status store.WebsiteStatus, responseTime int64) {
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	start := time.Now()

	res, err := client.Head(url)
	if err != nil {
		status = store.Down
		responseTime = time.Since(start).Milliseconds()
		return
	}

	switch {
	case res.StatusCode >= 200 && res.StatusCode <= 403:
		status = store.Up
	case res.StatusCode >= 500 && res.StatusCode <= 599:
		status = store.Down
	default:
		status = store.Unknown
	}

	responseTime = time.Since(start).Milliseconds()

	return status, responseTime
}

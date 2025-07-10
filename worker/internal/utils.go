package internal

import (
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/DevanshBhavsar3/echo/common/db/store"
)

func Ping(url string) (status store.WebsiteStatus, responseTime int64) {
	req, _ := http.NewRequest("HEAD", url, nil)
	req.Header.Set("User-Agent", "Echo-Monitor/1.0")

	trace := &httptrace.ClientTrace{}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	start := time.Now()

	res, err := http.DefaultTransport.RoundTrip(req)
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

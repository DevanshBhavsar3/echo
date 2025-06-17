package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/DevanshBhavsar3/common/store"
)

type Analytics struct {
	Url            string
	ResponseTimeMS int64
	Status         store.WebsiteStatus
}

func NewAnalytics(url string) *Analytics {
	return &Analytics{
		Url: url,
	}
}

// TODO: Add request timeout
func (a *Analytics) Ping() {
	req, _ := http.NewRequest("HEAD", a.Url, nil)
	req.Header.Set("User-Agent", "Echo-Monitor/1.0")

	var start, connect time.Time

	trace := &httptrace.ClientTrace{
		ConnectStart: func(network, addr string) {
			connect = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			time.Since(connect)
		},
		GotFirstResponseByte: func() {
			fmt.Printf("Time from start to first byte: %v\n", time.Since(start))
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case res.StatusCode >= 200 && res.StatusCode <= 403:
		a.Status = store.Up
	case res.StatusCode >= 500 && res.StatusCode <= 599:
		a.Status = store.Down
	default:
		a.Status = store.Unknown
	}

	a.ResponseTimeMS = time.Since(start).Milliseconds()
}

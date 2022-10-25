package racer

import (
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) string {
	if measureResponseTime(url1) < measureResponseTime(url2) {
		return url1
	}

	return url2
}

func measureResponseTime(url string) time.Duration {
	refTime := time.Now()
	http.Get(url)
	return time.Since(refTime)
}

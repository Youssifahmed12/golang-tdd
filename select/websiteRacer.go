package WebsiteRacer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	aDuartion := measureResponseTime(a)
	bDuartion := measureResponseTime(b)

	if aDuartion > bDuartion {
		return b, nil
	}
	return a, nil
}

func measureResponseTime(a string) time.Duration {
	startA := time.Now()
	http.Get(a)
	aDuartion := time.Since(startA)
	return aDuartion
}

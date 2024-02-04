package sync_demon

import (
	"fmt"
	"net/http"
	"time"
)

func ConfigurableRaceWebsites(siteA, siteB string, delay time.Duration) (string, error) {
	select {
	case <-ping(siteA):
		return siteA, nil
	case <-ping(siteB):
		return siteB, nil
	case <-time.After(delay):
		return "", fmt.Errorf("timed out waiting for %s and %s", siteA, siteB)
	}
}

func ping(url string) chan struct{} {
	tempChannel := make(chan struct{})
	go func() {
		http.Get(url)
		close(tempChannel)
	}()

	return tempChannel
}

func DefaultRaceWebsites(siteA, siteB string) (string, error) {
	return ConfigurableRaceWebsites(siteA, siteB, 10*time.Second)
}

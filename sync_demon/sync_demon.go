package sync_demon

import (
	"net/http"
	"time"
)

type websiteDuration struct {
	string
	duration time.Duration
}

func RaceWebsites(siteA, siteB string) string {

	durationChannel := make(chan websiteDuration)

	go GetWebsiteDuration(siteA, durationChannel)
	go GetWebsiteDuration(siteB, durationChannel)

	durations := make(map[string]time.Duration)

	dur := <-durationChannel
	durations[dur.string] = dur.duration
	dur = <-durationChannel
	durations[dur.string] = dur.duration

	if durations[siteA] < durations[siteB] {
		return siteA
	} else {
		return siteB
	}
}

func GetWebsiteDuration(url string, durChan chan (websiteDuration)) {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	durChan <- websiteDuration{url, duration}
}

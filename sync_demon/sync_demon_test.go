package sync_demon

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRaceWebsites(t *testing.T) {
	slowServer := GetDelayedTestServer(10 * time.Millisecond)
	fastServer := GetDelayedTestServer(0 * time.Millisecond)

	expected := fastServer.URL
	actual := RaceWebsites(slowServer.URL, fastServer.URL)

	defer slowServer.Close()
	defer fastServer.Close()

	if expected != actual {
		t.Errorf("expected: %s, but actual: %s", expected, actual)
	}
}

func GetDelayedTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(delay)
		res.WriteHeader(http.StatusOK)
	}))
}

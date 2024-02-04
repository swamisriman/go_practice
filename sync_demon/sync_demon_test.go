package sync_demon

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRaceWebsites(t *testing.T) {
	t.Run("Durations checker", func(t *testing.T) {
		slowServer := GetDelayedTestServer(20 * time.Millisecond)
		fastServer := GetDelayedTestServer(0 * time.Millisecond)

		expected := fastServer.URL
		actual, err := ConfigurableRaceWebsites(slowServer.URL, fastServer.URL, 1*time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		if err != nil {
			t.Error("expected: no time out, but actual: timed out")
		}
		if expected != actual {
			t.Errorf("expected: %s, but actual: %s", expected, actual)
		}
	})

	t.Run("Timeout checker", func(t *testing.T) {
		slowServer := GetDelayedTestServer(3 * time.Second)
		fastServer := GetDelayedTestServer(2 * time.Second)

		actual, err := ConfigurableRaceWebsites(slowServer.URL, fastServer.URL, 2*time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		if err == nil {
			t.Errorf("expected: time out, but actual: not timed out")
		}
		if actual != "" {
			t.Errorf("expected: %q, but actual: %q", "", actual)
		}
	})
}

func GetDelayedTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(delay)
		res.WriteHeader(http.StatusOK)
	}))
}

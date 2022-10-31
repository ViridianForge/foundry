package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("run the standard test", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)

		fastServer := makeDelayedServer(1 * time.Microsecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Error("got an error when none was expected")
		}

		if expected != got {
			t.Errorf("got %q, expected %q", got, expected)
		}

		slowServer.Close()
		fastServer.Close()
	})

	t.Run("return an error if a server doesn't respond by timeout", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Millisecond)
		serverB := makeDelayedServer(12 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error, but got none")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

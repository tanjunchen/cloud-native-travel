package selectortest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"
	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestRacer2(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer3(t *testing.T) {
	t.Run("", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	t.Run("returns an error if a server doesn't respond within 10s",func(t *testing.T){
		slowServer := makeDelayedServer(12 * time.Millisecond)
		fastServer := makeDelayedServer(11 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer2(slowServer.URL, fastServer.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
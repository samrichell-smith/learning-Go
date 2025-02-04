package selectpack

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(URL1, URL2 string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(URL1):
		return URL1, nil
	case <-ping(URL2):
		return URL2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", URL1, URL2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

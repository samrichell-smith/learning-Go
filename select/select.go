package selectpack

import (
	"net/http"
)

func Racer(URL1, URL2 string) (winner string) {
	select {
	case <-ping(URL1):
		return URL1
	case <-ping(URL2):
		return URL2
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

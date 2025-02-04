package selectpack

import (
	"net/http"
	"time"
)

func Racer(URL1, URL2 string) (winner string) {
	start1 := time.Now()
	http.Get(URL1)
	duration1 := time.Since(start1)

	start2 := time.Now()
	http.Get(URL2)
	duration2 := time.Since(start2)

	if duration1 < duration2 {
		return URL1
	}

	return URL2
}

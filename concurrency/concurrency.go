package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// makes a map of strings, the website names, as keys to a boolean value as to if it is visited
// creates a channel that takes structs of type result, which we have previously created
// for each url, create a goroutine with an anonymous function
// in each goroutine, we send a struct of type result into the channel we have made, with the given props to serve our results

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
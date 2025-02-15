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
// the goroutines are asynchronus, so we move to the next code will these run. then, we iterate for the number of results we have, and just take whatever the object in the channel is at the time, of which only one can go so we get one for each.

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
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	var mu sync.Mutex
	var wwg sync.WaitGroup
	cache := make(map[string]bool)

	var crawl func(string, int)

	crawl = func(url string, depth int) {
		defer wwg.Done()

		mu.Lock()
		if cache[url] {
			mu.Unlock()
			return
		}

		cache[url] = true
		mu.Unlock()

		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			wwg.Add(1)
			go crawl(u, depth-1)
			
		}

	}
	wwg.Add(1)
	go crawl(url, depth)
	wwg.Wait()

}

func main() {

	Crawl("https://golang.org/", 4, fetchermap)

}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetchermap = fakeFetcher{
	"https://golang.org/":         &a1,
	"https://golang.org/pkg/":     &a2,
	"https://golang.org/pkg/fmt/": &a3,
	"https://golang.org/pkg/os/":  &a4,
}

var a1 = fakeResult{
	body: "The Go Programming Language",
	urls: []string{
		"https://golang.org/pkg/",
		"https://golang.org/cmd/",
	},
}

var a2 = fakeResult{
	body: "Packages",
	urls: []string{
		"https://golang.org/",
		"https://golang.org/cmd/",
		"https://golang.org/pkg/fmt/",
		"https://golang.org/pkg/os/",
	},
}

var a3 = fakeResult{
	body: "Package fmt",
	urls: []string{
		"https://golang.org/",
		"https://golang.org/pkg/",
	},
}

var a4 = fakeResult{
	body: "Package os",
	urls: []string{
		"https://golang.org/",
		"https://golang.org/pkg/",
	},
}

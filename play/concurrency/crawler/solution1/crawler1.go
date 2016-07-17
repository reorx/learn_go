package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Fetcher interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// CrawlResult to contain
type CrawlResult struct {
	url  string
	body string
	urls []string
	err  error
}

// Cache to store
type Cache struct {
	c     map[string]CrawlResult
	ch    chan bool
	mutex sync.Mutex
}

// Add result to cache
func (c *Cache) Add(cr CrawlResult) {
	c.mutex.Lock()
	c.c[cr.url] = cr
	c.mutex.Unlock()
}

// Fetched to determine
func (c *Cache) Fetched(url string) bool {
	c.mutex.Lock()
	_, ok := c.c[url]
	defer c.mutex.Unlock()
	if ok {
		return true
	}
	return false
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *Cache) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	// determine if fetched
	if cache.Fetched(url) {
		log.Println("url is fetched", url)
		return
	}

	body, urls, err := fetcher.Fetch(url)
	// construct CrawlResult
	cr := CrawlResult{url, body, urls, err}
	cache.Add(cr)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		// goroutine
		go Crawl(u, depth-1, fetcher, cache)
	}
	return
}

func main() {
	cache := Cache{c: make(map[string]CrawlResult), ch: make(chan bool)}
	Crawl("http://golang.org/", 4, fetcher, &cache)

	time.Sleep(1 * time.Millisecond)
	fmt.Println("sleep 1")
	time.Sleep(1 * time.Second)
	fmt.Println("sleep done")
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
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

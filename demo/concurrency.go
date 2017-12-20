package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

type URLFetcher interface {
	FetchURL(url string) (*http.Response, error)
}

type HTTPFetcher struct{}

func (f HTTPFetcher) FetchURL(url string) (*http.Response, error) {
	return http.Get(url)
}

type ResponseInfo struct {
	resp        *http.Response
	err         error
	elapsedTime time.Duration
}

func fetchURLs(fetcher URLFetcher, urls ...string) <-chan ResponseInfo {
	c := make(chan ResponseInfo)

	go func() {
		wg := sync.WaitGroup{}
		for _, u := range urls {
			wg.Add(1)

			go func(u string) {
				defer wg.Done()

				start := time.Now()
				resp, err := fetcher.FetchURL(u)
				c <- ResponseInfo{
					resp:        resp,
					err:         err,
					elapsedTime: time.Since(start),
				}
			}(u)
		}
		wg.Wait()
		close(c)
	}()

	return c
}

func main() {
	start := time.Now()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	for info := range fetchURLs(HTTPFetcher{}, os.Args[1:]...) {
		if info.err != nil {
			fmt.Fprintln(w, "error:", info.err)
		} else {
			fmt.Fprintf(w, "%s\t%s\n", info.elapsedTime, info.resp.Request.URL)
		}
	}

	fmt.Fprintf(w, "%s\ttotal time\n", time.Since(start))
	w.Flush()
}

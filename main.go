package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mazahireyvazli/go-myhttp/cmd"
)

var parallelFlag = flag.Int("parallel", 10, "number of parallel requests")

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: go-myhttp [flags] URL1 [URL2 ...]")
		fmt.Fprintln(os.Stderr, "flags:")
		flag.PrintDefaults()
	}
	flag.Parse()

	urls := flag.Args()
	parallel := *parallelFlag

	if len(urls) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	numWorkers := len(urls)
	if parallel < numWorkers {
		numWorkers = parallel
	}

	httpClient := cmd.NewMyHTTPClient()
	responsech := httpClient.CreateWorkers(numWorkers, urls)

	for i := 0; i < len(urls); i++ {
		response := <-responsech
		if response != nil {
			println(response.URL, response.ResponseBodyHash)
		}
	}
}

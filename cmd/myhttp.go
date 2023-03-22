package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"net/url"
	"time"
)

type MyHTTPResponse struct {
	URL              string
	ResponseBodyHash string
}

type MyHTTPClient struct {
	client http.Client
}

func (c *MyHTTPClient) parseRequestURL(u string) (*url.URL, error) {
	parsedURL, err := url.Parse(u)

	if err != nil {
		return nil, err
	}

	// set default scheme
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}

	return parsedURL, nil
}

func (c *MyHTTPClient) hashResponseBody(body []byte) string {
	hash := md5.Sum(body)
	return hex.EncodeToString(hash[:])
}

func (c *MyHTTPClient) SendHTTPRequest(requestURL string) (*MyHTTPResponse, error) {
	parsedURL, err := c.parseRequestURL(requestURL)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Get(parsedURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &MyHTTPResponse{
		URL:              parsedURL.String(),
		ResponseBodyHash: c.hashResponseBody(respBody),
	}, nil
}

func (c *MyHTTPClient) worker(requestch <-chan string, responsech chan<- *MyHTTPResponse) {
	for url := range requestch {
		// Please don't add any extra features to the tool.
		// so skip error handling
		resp, _ := c.SendHTTPRequest(url)

		responsech <- resp
	}
}

func (c *MyHTTPClient) CreateWorkers(num int, urls []string) chan *MyHTTPResponse {
	requestch := make(chan string, len(urls))
	responsech := make(chan *MyHTTPResponse, len(urls))

	for i := 0; i < num; i++ {
		go c.worker(requestch, responsech)
	}

	for _, url := range urls {
		requestch <- url
	}
	close(requestch)

	return responsech
}

func NewMyHTTPClient() *MyHTTPClient {
	return &MyHTTPClient{
		// net/http clients are safe for concurrent use by multiple goroutines.
		client: http.Client{
			Timeout: time.Second * 5,
		},
	}
}

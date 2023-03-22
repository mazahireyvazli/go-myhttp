# MyHTTP Tool

MyHTTP is a simple command-line tool for sending HTTP GET requests to multiple URLs in parallel and computing the MD5 hash of the response body. The tool supports specifying the number of parallel requests to be made.

## Installation

To install MyHTTP, simply clone this repository and run the following command in the root directory:

`go install github.com/mazahireyvazli/go-myhttp@latest`

This will install the `go-myhttp` binary in your Go bin directory.

## Usage

To use MyHTTP, run the `go-myhttp` command followed by the URLs you want to request:

`go-myhttp [flags] [URLs...]`

The following flags are available:

- `-parallel`: the number of parallel requests to make (default 10)

The tool will send HTTP GET requests to the specified URLs in parallel, compute the MD5 hash of the response body, and print the URL and hash for each successful request.

## Test

To run tests, use the following command:

`make test`

## Build

To build the `go-myhttp` binary, use the following command:

`make build`

This will generate the `go-myhttp` binary in the `bin` directory.

## Example

To make 2 parallel requests to example.com and google.com, run the following command:

`go-myhttp -parallel 2 https://example.com https://google.com`

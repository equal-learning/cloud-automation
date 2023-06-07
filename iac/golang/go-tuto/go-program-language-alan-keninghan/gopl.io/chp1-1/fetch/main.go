// Fetchall fetches URLs in parallel and reports their times and sizes.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	recv_results := make(chan string)
	if len(os.Args) == 1 {
		os.Exit(1)
	} else {

		for _, url := range os.Args[1:] {

			go fetch(url, recv_results)
		}

		for range os.Args {

			str := <-recv_results

			fmt.Printf("%s", str)

		}

	}

	//recv_results := make(chan string)
	//fetch("https://google.com")
	//fetch("https://harvard.edu")

}

func fetch(urlpath string, ch chan string) {

	time_starting := time.Now()

	resp, err := http.Get(urlpath)

	if err != nil {
		fmt.Printf("Error reading urls")
		fmt.Fprintf(os.Stderr, "error reading %v", err)
		os.Exit(1)

	} else {
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Print("Error reading body")
			fmt.Fprintf(os.Stderr, "error reading %v", err)
			ch <- fmt.Sprintf("error reading %v", err)
			os.Exit(1)
		} else {

			duration := time.Since(time_starting).Seconds()
			size := len(body)

			fmt.Printf("size : %d", len(body))
			fmt.Printf("%f %d %s", duration, size, urlpath)
			// ch <- fmt.Sprintf("%f %d %s", duration, size, urlpath)

		}

	}

}

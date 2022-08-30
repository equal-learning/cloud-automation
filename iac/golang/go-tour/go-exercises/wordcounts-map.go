/*

Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	substrings := strings.Fields(s)
	wdcounts := make(map[string]int)

	for _, substring := range substrings {
		if _, ok := wdcounts[substring]; ok {
			wdcounts[substring] += 1

		} else {
			wdcounts[substring] = 1
		}
	}

	return wdcounts
}

func main() {
	wc.Test(WordCount)
}

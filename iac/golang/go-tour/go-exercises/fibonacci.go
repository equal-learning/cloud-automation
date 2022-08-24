package main

import "fmt"

/*
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	n := 0
	nMinus1 := 0
	nMinus2 := 0

	return func() int {

		if n == 0 {
			n++

			return 0

		} else if n == 1 {
			n++
			nMinus1 = 1
			return 1
		} else {
			res := nMinus1 + nMinus2
			nMinus2 = nMinus1
			nMinus1 = res
			return res
		}

	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
DEFER stacking ::

Deferred function calls are pushed onto a stack.
When a function returns, its deferred calls are executed in last-in-first-out (LIFO) order.

For more on defer : https://go.dev/blog/defer-panic-and-recover

 */

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
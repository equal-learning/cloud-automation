/*

A function can take zero or more arguments.
More on Go's declaration syntax : https://go.dev/blog/declaration-syntax
*/

package main

import "fmt"

func add1(x int, y int) int {
	return x + y
}

// Syntactic sugar : When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add1(42, 13))
}

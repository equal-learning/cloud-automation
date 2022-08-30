/*

In addition to generic functions, Go also supports generic types.
A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.

This example demonstrates a simple type declaration for a singly-linked list holding any type of value.

As an exercise, add some functionality to this list implementation.

*/

package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}

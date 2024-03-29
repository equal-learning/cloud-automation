/*

The empty interface ::


The interface type that specifies zero methods is known as the empty interface:

interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of "unknown type".
For example, fmt.Print takes any number of arguments of type interface{}.

*/

package main

import "fmt"

func main() {
	var i interface{} // Remark :: Empty interface :: (<nil>, <nil>)
	describe(i)

	i = 42
	describe(i) // Remark :: (42, int)

	i = "hello"
	describe(i) // Remark :: (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

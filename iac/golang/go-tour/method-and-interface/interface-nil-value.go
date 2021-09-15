/*

Nil interface values ::

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

 */

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()       // Remark :: (<nil>, <nil>) ::  "Throws an error" as there is no type inside the interface
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

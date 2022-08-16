/*

Interface values ::

Under the hood, interface values can be thought of as a "tuple" of a value and a concrete type: (value, type)

An interface value holds a "value of a specific underlying concrete type".

Calling a method on an interface value executes the method of the same name on its underlying type.

*/

package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
	N()
}

type T struct {
	S_M string
	S_N string
}

func (t *T) M() { // Remark :: Implicitely implementing interface I
	fmt.Println(t.S_M)
}

func (t *T) N() { // Remark :: Implicitely implementing interface I
	fmt.Println(t.S_N)
}

type F float64

func (f F) M() { // Remark :: Implicitely implementing interface I
	fmt.Println(f)
}

func (f F) N() { // Remark :: Implicitely implementing interface I
	fmt.Println(f)
}

func main() {
	var i I // Remark :: Polymorphisme

	i = &T{"Hello", "world"}
	describe(i)
	i.M()
	i.N()

	i = F(math.Pi)
	describe(i)
	i.M()
	i.N()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

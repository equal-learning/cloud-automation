/*

Interfaces ::

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 22.
Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).

 */

package main

import (
	"fmt"
	"math"
)

type Abser interface {    // Remark : An interface type is defined as a set of method signatures.
	Abs() float64
}

func main() {
	var a Abser    // :: Remark :: A value of interface type can hold any value that implements its methods.
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f              // Remark :: a MyFloat implements Abser
	a = &v            // Remark :: a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v       // Remark :: It is an error !!!!

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

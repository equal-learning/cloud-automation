/*

FUNCTION CLOSURE ::

Go functions may be closures.
A closure is a function value that references variables from **outside its body**.
The function may access and assign to the referenced variables; in this sense the function is "BOUND" to the variables.

For example, the adder function returns a closure. Each closure is bound to its **OWN** sum variable.

Remark :: Easy to implement the function "that can recall"
*/

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int { // returing a closure (which lookes like first class functions !!)
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

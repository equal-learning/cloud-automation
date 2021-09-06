/*

POINTERS ::

Go has pointers. A pointer holds the memory address of a value.
The type *T is a pointer to a T value. Its zero value is nil.

Unlike C, Go has no pointer arithmetic. Happy days :)

 */

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i, The & (address of) operator generates a pointer to its operand.
	fmt.Println(*p) // read i through the pointer, The * (indirection) operator denotes the pointer's underlying value
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

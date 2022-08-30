/*

"Mutating Maps" ::

m := make(map[string]int) // creating an empty map ready to use

Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Note: If elem or ok have not yet been declared you could use a short declaration form:

elem, ok := m[key]


*/

package main

import "fmt"

func main() {
	m := make(map[string]int) // make function returns a map of the given type, initialized and ready for use

	m["Answer"] = 42 // insert
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48 // update
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // delete
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // test if an element is present
	fmt.Println("The value:", v, "Present?", ok)
}

package main

import "fmt"

func main() {

	org := []int{1, 2, 3, 4, 5}
	rev := make([]int, len(org))

	for i := 0; i < len(org); i++ {
		rev[i] = org[len(org)-(i+1)]
	}

	fmt.Println(org)
	fmt.Println(rev)

	// using swaping
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	for i, j := 0, len(org)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]

	}
	fmt.Println(s)

}

package main

import "fmt"

// Package treesort provides insertion sort using an unbalanced binary tree.

func main() {

	var elms []int = []int{5, 3, 4, 9, 2, 6}
	var sortd []int = make([]int, 0, len(elms))

	fmt.Println(elms)
outer:
	for _, elm := range elms {

		for i := 0; i < len(sortd); i++ {
			if elm <= sortd[i] {
				sortd = append(sortd, 0)
				insert(sortd, elm, i)
				continue outer
			}
		}
		sortd = append(sortd, elm)
		fmt.Println(sortd)
	}

	fmt.Println(sortd)

}

func insert(arr []int, elm, pos int) []int {

	for j := len(arr) - 1; pos < j; j-- {
		arr[j] = arr[j-1]

	}
	arr[pos] = elm
	return arr
}

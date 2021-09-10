/*

FUNCTION CLOSURE example ::

Implementing fibonacci series using " function closure".

 */

 package main

import "fmt"

func fibNaive(n int) int {

	 if n == 0 {
		 return 0
	 } else if n == 1 {
		return 1
	} else {
		 return fibNaive(n-1) + fibNaive(n-2)
	 }

 }

func fibSeries(n int) func (int)[]int  {

	elm := []int{1,1}


	return func (num int) []int  {

		for i:=2;i<n+1;i++ {
			elm = append(elm, elm[i-1] + elm[i-2])
		}

		return elm

	}

	}



func main () {

     fibnumbers := fibSeries(100)(100)

     for i,v := range fibnumbers {

     	fmt.Println(i+1,v)

	}

 }
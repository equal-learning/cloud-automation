/*

Select ::

The select provides "a control structure which give a way to contorl access from multiple channels at once".
The select statement lets a goroutine "wait on multiple communication operations".

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.


*/

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // remark: Which channel to select ?? At each iteration one of the ready conditions selected randomly is executed
		case c <- x: // remark :: c is read till 9th fibonacci number
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit") // remark :: waits till the 9th fibonacci number
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

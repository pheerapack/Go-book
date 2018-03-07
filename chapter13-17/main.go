package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
		//close(naturals)
	}()
	for {
		fmt.Println(<-squares)
	}
}

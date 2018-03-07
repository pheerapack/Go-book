package main

import "fmt"

func main() {
	c := make(chan int, 3) // innitial buffer number
	c <- 1                 // input buffer can be less than 4,max 3
	c <- 2
	c <- 3 // overfilled the buffer
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

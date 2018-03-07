package main

import (
	"fmt"
)

func main() {
	array := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)             //create channel
	go sum(array[:len(array)/2], ch) // do 0 to 3 sum
	go sum(array[len(array)/2:], ch) // do 4 to 6 sum
	x, y := <-ch, <-ch               // input to x , y
	fmt.Println(x, y, x+y)
}
func sum(array []int, ch chan int) {
	sum := 0
	for _, value := range array {
		sum += value
	}
	ch <- sum
}

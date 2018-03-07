package main

import "fmt"

func main() {
	ch := make(chan int, 26) // buffer , number = number that can recive
	go func() { ch <- 1 }()
	ch <- 5
	ch <- 5
	ch <- 5
	ch <- 5
	fmt.Println("cap:", cap(ch))
	fmt.Println("len:", len(ch))
}

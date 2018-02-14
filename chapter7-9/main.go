package main

import "fmt"

func main() {
	slices := [5]int{1,2,3,4,5}
	for i,number := range slices {
		fmt.Println(i,number)
	}
}
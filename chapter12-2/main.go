package main

import "fmt"

func main() {
	//decare fmt and input in function
	fmt.Println(
		func(a, b int) int {
			return a + b
		}(2, 3))
}

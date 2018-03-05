package main

import "fmt"

func main() {
	addfunc := func(a int) func(b int) int {
		return func(b int) int {
			return a + b
		}
	}
	add2 := addfunc(2)
	fmt.Println(add2(3))

	add3 := addfunc(25)
	fmt.Println(add3(3))

}

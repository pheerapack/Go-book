package main

import "fmt"

func main() {
	var addVar func(int, int) int //decare input int,int return int in variable
	addVar = func(a, b int) int { //add func behavior
		return a + b
	}
	fmt.Println(addVar(2, 3)) //input
}

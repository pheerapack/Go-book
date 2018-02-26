package main

import (
	"fmt"
)

type shape interface {
	area() int
}

type traingle struct {
	height int
	base   int
}

func (a traingle) area() int {
	return a.height * a.base / 2
}

func main() {
	var p1 shape = traingle{height: 3, base: 4}
	fmt.Println(p1.area())
}

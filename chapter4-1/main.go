package main

import "fmt"

var a string

func main(){
	//Long declaration
	var x string = "Hello World Long declaration"
	fmt.Println(x)
	var y string
	y="Hello World Long declaration"
	fmt.Println(y)

	//Short declaration
	z:="Hello World Short declaration"
	fmt.Println(z)
	fmt.Printf("Type : %T\n", z)

	f()
}
func f(){
	a="a is Global"
	fmt.Println(a)

	const b string = "b is const , Hello World"
	fmt.Println(b)
	var (
		c=1
		e=2
		f=3
	)
	fmt.Println(c,e,f)

	v1,v2 := "First","2"
	fmt.Println(v1,v2)
}
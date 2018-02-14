package main

import "fmt"

func main()  {
	for j := 1; j <= 100; j++ {
		r := findfizzbuzz(j)
		fmt.Println(j,r)
	}
}

func findfizzbuzz(j int) (r string) {
	if j%15 == 0 {
		r = "FizzBuzz"
	} else if j%3 == 0 {
		r = "Fizz"
	} else if j%5 == 0 {
		r = "Buzz"
	} else {
		r = "Number"
	}
	return r
}
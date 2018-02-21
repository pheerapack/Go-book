package main

import (
	"fmt"
	"strconv"
)

func main() {
	forloop()
}

func forloop() {
	for j := 1; j <= 100; j++ {
		r := findfizzbuzz(j)
		fmt.Println(j, r)
	}
}

func findfizzbuzz(j int) (r string) {
	mod := [3]int{15, 3, 5}
	answer := [3]string{"FizzBuzz", "Fizz", "Buzz"}
	for i := 0; i < len(mod); i++ {
		if j%mod[i] == 0 {
			return answer[i]
		}
	}
	return strconv.Itoa(j)
}

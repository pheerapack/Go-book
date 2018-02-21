package main

import "fmt"

func main() {
	mapt := map[int]int{1: 1, 2: 2, 3: 3}
	double(mapt)
	fmt.Printf("origin addr %p\n", mapt)
	fmt.Printf("original %v\n", mapt)
}
func double(nums map[int]int) {
	fmt.Printf("double addr %p\n", nums)
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(nums)
}

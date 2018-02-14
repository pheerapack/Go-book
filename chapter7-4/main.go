package main
import "fmt"
func main() {
	slice := []int{1,2,3}
	fmt.Println(slice)
	newslice:=make([]int,2)
	fmt.Println(newslice)

	//copy way : move newslice to slice
	copy(slice,newslice)
	fmt.Printf("slice: %v\n",slice)
	fmt.Printf("slice: %v\n",newslice)
}
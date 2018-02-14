package main
import "fmt"
func main() {
	x, y := f()
	fmt.Println(x,y)
}


func f() (int, int) {
	return 8, 9
}

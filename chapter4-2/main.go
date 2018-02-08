package main
import "fmt"
func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (( 5 *(input - 32.0)) / 9.0)
	//a := float64(output)
	fmt.Printf("%.2f",output)
	

	//i := 5
   // f := float64(i)
   // fmt.Printf("f is %f\n", f)
}
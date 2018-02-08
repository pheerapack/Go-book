package main
import "fmt"
func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (( 5 *(input - 32.0)) / 9.0)
	//a= input
	fmt.Println(input)
	//fmt.Sprintf("%.2f",input)
	fmt.Printf("Fahrenheit into Celsius : %.2f",output)
	fmt.Sprintf("%.6f", 21312421.213123)
	
	//i := 5
   // f := float64(i)
   // fmt.Printf("f is %f\n", f)
}
package main

import "fmt"

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few clond"))
	fmt.Println(weatherCelsius(34, "Tak sunny"))
	fmt.Println(weatherCelsius(17, "Phuket rainny"))
	fmt.Println(weatherCelsius(9, "Chiangmai cold"))

}

func weatherCelsius(cel int, province string) (r string, province2 string) {
	if cel == 25 {
		r = "---  ---\n  |  |\n---  ---\n|       |\n---  --- 'c"
	}
	if cel == 34 {
		r = "---      \n   |  |   |\n---  ---\n    |     |\n---      'c"
	}
	if cel == 17 {
		r = "---  ---\n  |  |\n---  ---\n|       |\n---  --- 'c"
	}
	if cel == 9 {
		r = "---  ---\n  |  |\n---  ---\n|       |\n---  --- 'c"
	}
	r = r + "\n" + province
	return "", r
}

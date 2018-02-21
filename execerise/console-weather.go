package main

import "fmt"

func maiewfewwn() {
	fmt.Println(weatherCelsius(25, "Bangkok few clond"))
	//fmt.Println(weatherCelsius(34, "Tak sunny"))
	//fmt.Println(weatherCelsius(17, "Phuket rainny"))
	//fmt.Println(weatherCelsius(9, "Chiangmai cold"))

}

func weatherCelsius(cel int, province string) (r string) {
	maps := map[int]string{
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
		0: "0",
	}

	for key, number := range maps {
		fmt.Println(key, number)
	}

	if cel == 25 {
		r = "---  ---\n  |  |\n---  ---\n|       |\n---  --- 'c"
	}

	r = r + "\n" + province
	return r
}

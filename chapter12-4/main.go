package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
			var addVar func(i int) string //decare input int,int return int in variable
			addVar = func(i int) string { //add func behavior
				if i%15 == 0 {
					return strconv.Itoa(i) + " FizzBuzz"
				} else if i%3 == 0 {
					return strconv.Itoa(i) + " Fizz"
				} else if i%5 == 0 {
					return strconv.Itoa(i) + " Buzz"
				} else {
					return strconv.Itoa(i)
				}
			}
			for i := 1; i <= 100; i++ {
				if i%15 == 0 {
					fmt.Println(addVar(i)) //input
				}
			}

			/*
				for i := 1; i <= 100; i++ {
					fmt.Println(
						func(i int) string {
							if i%15 == 0 {
								return strconv.Itoa(i) + " FizzBuzz"
							} else if i%3 == 0 {
								return strconv.Itoa(i) + " Fizz"
							} else if i%5 == 0 {
								return strconv.Itoa(i) + " Buzz"
							} else {
								return strconv.Itoa(i)
							}

						}(i))
				}


		addfunc := func(a, b string) (string, string) {
			return a, b
		}

		for i := 1; i <= 100; i++ {
			if i%15 == 0 {
				fmt.Println(addfunc(strconv.Itoa(i), "Fizzbuzz"))
			} else if i%3 == 0 {
				fmt.Println(addfunc(strconv.Itoa(i), "Fizz"))
			} else if i%5 == 0 {
				fmt.Println(addfunc(strconv.Itoa(i), "buzz"))
			} else {
				fmt.Println(addfunc(strconv.Itoa(i), strconv.Itoa(i)))
				//return strconv.Itoa(i)
			}
		}
	*/
	for i := 1; i <= 100; i++ {
		fmt.Println(i, findfizzbuzz(i))
	}

}

func fbtemplate(fbnumber int, str string) func(int) (string, bool) {
	return func(n int) (string, bool) {
		if n%fbnumber == 0 {
			return str, true
		}
		return "", false
	}
}

func findfizzbuzz(number int) string {
	//Array part , call function
	Arrayfb := [...]func(n int) (string, bool){
		fbtemplate(3, "Fizz"),
		fbtemplate(5, "buzz"),
		fbtemplate(15, "Fizzbuzz"),
	}

	//For loop find function if return true > function
	for i := 0; i < len(Arrayfb); i++ {
		if str, ok := Arrayfb[i](number); ok {
			return str
		}
	}
	return strconv.Itoa(number)

}

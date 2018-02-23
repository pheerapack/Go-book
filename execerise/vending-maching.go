package main

import "fmt"

type NewVendingMachine struct {
	coins string
	items string
}

func (c NewVendingMachine) InsertCoin() string {

	if c.coins == "T" {
		r := 10
	}
	return r
}

func main() {
	vm := NewVendingMachine{"T", "SD"}
	fmt.Println(vm.InsertCoin())
	//vm.InsertCoin("F")
	//vm.InsertCoin("TW")
	//vm.InsertCoin("O")
	//fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18
	//can := vm.SelectSD()
	//fmt.Println(can) // SD
}

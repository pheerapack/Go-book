package main

import "fmt"
import "vendingmachine"

var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
var items = map[string]int{"SD": 18, "CC": 12}

func main() {

	vm := vendingmachine.NewVendingMachine(coins, items)
	fmt.Println("Insert Money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Insert Money:", vm.InsertedMoney())
	can := vm.SelectSD()
	fmt.Println(can)

	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Insert Money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can)

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Insert Money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can)

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("TW")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	coin := vm.CoinsReturn()
	fmt.Println(coin)
}

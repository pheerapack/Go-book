package main

import "fmt"

type VendingMachine struct {
	insertedMoney int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney = 10
}

/*
func (c NewVendingMachine) InsertCoin() string {
	maps := map[string]string{
		"T": "10",
		"F": "5",
		"O": "1",
	}
	return maps[c.coins]
}
*/
func main() {
	//var p1 InsertCoin = NewVendingMachine{10}
	//var p2 InsertCoin = NewVendingMachine{"F"}
	//var p3 InsertCoin = NewVendingMachine{"TW"}
	//var p4 InsertCoin = NewVendingMachine{"O"}
	//fmt.Println(p1.GetInsertedMoney())

	vm := VendingMachine{}
	fmt.Println("Insert Money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	fmt.Println("Insert Money:", vm.InsertedMoney())
	//vm.InsertCoin("T")
	//vm.InsertCoin("F")
	//vm.InsertCoin("TW")
	//vm.InsertCoin("O")
	//fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18

	//can := vm.SelectSD()
	//fmt.Println(can) // SD
}

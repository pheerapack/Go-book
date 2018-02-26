package main

import "fmt"

type VendingMachine struct {
	insertedMoney int
	coins         map[string]int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
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
	vm := VendingMachine{}
	fmt.Println("Insert Money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Insert Money:", vm.InsertedMoney())
	//vm.InsertCoin("T")
	//vm.InsertCoin("F")
	//vm.InsertCoin("TW")
	//vm.InsertCoin("O")
	//fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18

	//can := vm.SelectSD()
	//fmt.Println(can) // SD
}

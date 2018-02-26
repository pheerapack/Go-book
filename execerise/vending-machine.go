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

func (m *VendingMachine) SelectSD() string {
	m.insertedMoney = 0
	return "SD"
}
func (m *VendingMachine) SelectCC() string {
	price := 12
	change := m.insertedMoney - price
	return "CC" + m.change(change)
}

func (m VendingMachine) change(c int) (change string) {
	return ",F,TW,O"
}
func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	vm := VendingMachine{coins: coins}
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
}

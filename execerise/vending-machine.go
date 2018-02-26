package main

import "fmt"

type VendingMachine struct {
	insertedMoney int
	coins         map[string]int
	items         map[string]int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine) SelectSD() string {
	//m.insertedMoney = 0
	price := m.items["SD"]
	change := m.insertedMoney - price
	m.insertedMoney = 0
	return "SD" + m.change(change)
}
func (m *VendingMachine) SelectCC() string {
	//m.insertedMoney = 0
	price := m.items["CC"]
	change := m.insertedMoney - price
	m.insertedMoney = 0
	return "CC" + m.change(change)
}

func (m *VendingMachine) CoinsReturn() string {
	coins := m.change(m.insertedMoney)
	m.insertedMoney = 0
	return coins[2:len(coins)]
}

func (m *VendingMachine) change(c int) string {
	var str string
	values := [...]int{10, 5, 2, 1}
	coins := [...]string{"T", "F", "TW", "O"}

	for i := 0; i < len(values); i++ {
		if c >= values[i] {
			str += "," + coins[i]
			c -= values[i]
			i--
			//Change 10,5,2,1 asc
		}

	}
	return str
}
func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	var items = map[string]int{"SD": 18, "CC": 12}
	vm := VendingMachine{coins: coins, items: items}
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
	fmt.Println("Inserted Money 5:", vm.InsertedMoney())
	coin := vm.CoinsReturn()
	fmt.Println(coin)
}

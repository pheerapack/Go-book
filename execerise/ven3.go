package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
/*
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

func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	vm := VendingMachine{coins: coins}
	fmt.Println("Insert Money:", vm.InsertedMoney())

	snr := bufio.NewScanner(os.Stdin)
	enter := "Enter a line of data:"
	for fmt.Println(enter); snr.Scan(); fmt.Println(enter) {
		line := snr.Text()
		if len(line) == 0 {
			break
		}
		fields := strings.Fields(line)
		fmt.Printf("Fields: %q\n", fields)

	}
	if err := snr.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
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

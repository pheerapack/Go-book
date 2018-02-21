package main

import (
	"fmt"
)

func main() {
	//const int lcd_segment[][COL] =
	lcd_segment := [][]uint8{
		{1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1}, // 0
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0}, // 1
		{1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1}, // 2
		{1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1}, // 3
		{0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0}, // 4
		{1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1}, // 5
		{1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1}, // 6
		{1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, // 7
		{1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1}, // 8
		{1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1}, // 9
	}

	//slices := [5]int{1,2,3,4,5}
	segment_c := [...]string{"-", "-", "|", "|", "|", "-", "-", "|", "|", "|", "-", "-"}
	segment_x := [12]int{1, 3, 0, 2, 4, 1, 3, 0, 2, 4, 1, 3}
	segment_y := [12]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4}
	//const char segment_c[] := "--|||--|||--"
	// const char segment_x[] = { 1, 3, 0, 2, 4, 1, 3, 0, 2, 4, 1, 3 }
	// const char segment_y[] = { 0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4 }

	var display [5][5]string
	var display2 [5][5]string
	for digit := 0; digit < 2; digit++ {
		display = display2
		//memset(display, ' ', sizeof(display))
		for segnum := 0; segnum < 12; segnum++ {
			if lcd_segment[digit][segnum] == 1 {
				x := segment_x[segnum]
				y := segment_y[segnum]
				display[y][x] = segment_c[segnum]
			}
		}
		for a := 0; a < 5; a++ {
			for b := 0; b < 5; b++ {
				fmt.Print(display[a][b])
				//putchar(display[y][x])
				fmt.Print("\n")
			}
		}
	}
}

package main

import "fmt"

func main() {
	col := 1
	row := 1
	var p1 int64 = 20151125

	for col != 3083 || row != 2978 {
		row = row - 1
		col = col + 1
		if row == 0 {
			row = col
			col = 1
		}
		p1 = (p1 * 252533) % 33554393
	}
	fmt.Println(row, col, p1)
}

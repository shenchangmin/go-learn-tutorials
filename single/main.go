package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	x := 3
	y := 4
	z := 5
	max_xy := max(x, y)
	max_xz := max(x, z)
	max_yz := max(y, z)
	fmt.Println(max_xy)
	fmt.Println(max_xz)
	fmt.Println(max_yz)
}

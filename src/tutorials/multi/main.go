package main

import "fmt"

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}
func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("and the number is:%d\n", n)
	}
}
func main() {
	myfunc(1, 2, 3, 4, 5)
	x := 3
	y := 4
	xPLUSy, xTIMESy := SumAndProduct(x, y)
	fmt.Printf("%d+%d=%d\n", x, y, xPLUSy)
	fmt.Printf("%d*%d=%d\n", x, y, xTIMESy)
}

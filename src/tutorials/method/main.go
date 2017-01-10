package main

import (
	"fmt"
	"math"
)

// Rectangle 长方形
type Rectangle struct {
	width, height float64
}

// Circle 圆
type Circle struct {
	radius float64
}

//计算长方形面积
func (r Rectangle) area() float64 {
	return r.width * r.height
}

//计算圆的面积
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

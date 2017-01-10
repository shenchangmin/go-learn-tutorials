package main

import "fmt"

const (
	// WHITE 颜色常量
	WHITE = iota
	// BLACK 颜色常量
	BLACK
	// BLUE 颜色常量
	BLUE
	// RED 颜色常量
	RED
	// YELLOW 颜色常量
	YELLOW
)

// Color 颜色
type Color byte

// Box 盒子
type Box struct {
	width, height, depth float64
	color                Color
}

// BoxList a slice of boxes
type BoxList []Box

// Volume 盒子的体积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// SetColor 设置盒子的颜色
func (b Box) SetColor(c Color) {
	b.color = c
}

// BiggestColor 体积最大的盒子的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

// PaintItBlack 将盒子颜色设置成黑色
func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}
func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}

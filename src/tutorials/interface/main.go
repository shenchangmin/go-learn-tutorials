package main

import (
	"fmt"
	"strconv"
)

//Element 存值
type Element interface {
}

//List Element
type List []Element

//Person class
type Person struct {
	name string
	age  int
}

//定义了string方法，实现了fmt.stringer
func (p Person) String() string {
	return "(name:" + p.name + "- age:" + strconv.Itoa(p.age) + "years)"
}
func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"Dennis", 70}
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

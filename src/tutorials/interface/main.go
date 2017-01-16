package main

import (
	"fmt"
	"strconv"
)

//Human 人基础类
type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "<<" + h.name + "-" + strconv.Itoa(h.age) + "years old-" + h.phone + ">"
}
func main() {
	Bob := Human{"Bob", 39, "000-7777-xx"}
	fmt.Println("This Human is : ", Bob)
}

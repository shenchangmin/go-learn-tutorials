package main

import "fmt"

// Human 基础类
type Human struct {
	name  string
	age   int
	phone string
}

// Student 学生类
type Student struct {
	Human
	school string
	loan   float32
}

// Employee 雇主类
type Employee struct {
	Human
	company string
	money   float32
}

// SayHi Human对象实现SayHi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

// Sing Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Guzzle Human的方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

//SayHi Employee对象重载SayHi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s.Call me on %s\n", e.name, e.company, e.phone)
}
func main() {

}

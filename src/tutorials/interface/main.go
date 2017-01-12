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

//BorrowMoney Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//SpendSalary Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

//Men 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

// YoungChap 年轻人活动
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

// ElderlyGent 老年人活动
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {

}

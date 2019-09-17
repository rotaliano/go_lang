package main

import (
	"fmt"
)

// Person
type Person struct{ Name string }
type MyFloat float64

// 値レシーバ
// Person 型に対してメソッドを定義する
func (p Person) Greet(msg string) {
	fmt.Printf("%s, I'm %s.\n", msg, p.Name)
}

// ポインタレシーバ
// Person 型に対してメソッドを定義する
func (pp *Person) Shout(msg string) {
	fmt.Printf("%s!!!\n", msg)
}

func (f MyFloat) Abs() float64 {
	if ( f < 0) {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	p := Person{Name: "Taro"}
	p.Greet("Hi")

	pp := &Person{Name: "Taro"}
	pp.Shout("OH MY GOD")
}

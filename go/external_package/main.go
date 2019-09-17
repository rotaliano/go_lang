package main

import (
	a "./foo"
)

func main() {
	// 同じパッケージの関数call
	// go run main.go human.go と実行
	// または
	// go run go/external_package
	mike := Person{Name: "Miki", Age: 11}
	mike.Say()

	// 別パッケージの関数call
	a.Hello()
}

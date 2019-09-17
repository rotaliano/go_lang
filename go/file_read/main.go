package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("aaa.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	fmt.Println(string(b))
}

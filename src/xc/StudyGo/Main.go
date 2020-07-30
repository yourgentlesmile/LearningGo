package main

import "fmt"

//单独声明变量
var single string

//批量声明变量
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "new World"
	age = 18
	isOk = true
	fmt.Print("Hello world")
}

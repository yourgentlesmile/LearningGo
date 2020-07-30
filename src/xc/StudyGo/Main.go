package main

import (
	"fmt"
	"unicode"
)

//单独声明变量
var single string

//批量声明变量
var (
	name string
	age  int
	isOk bool
)

func foo() (int, string) {
	return 10, " "
}

func varDeclare() {
	name = "new World"
	age = 18
	isOk = true
	//非全局变量声明后必须使用
	var hhh string = "123"
	//类型推断
	var bbb = "123"
	//简短变量声明，只能在函数里面用
	ccc := "456"
	fmt.Print("Hello world" + hhh + bbb + ccc)

	//匿名变量，在使用多重赋值时，如果想要忽略某个值
	//可以使用匿名变量。匿名变量用一个下划线表示
	//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	x, _ := foo()
	_, y := foo()
	fmt.Print(x)
	fmt.Print(y)

	//定义常量的时，可以单独声明常量
	const cc1 = "const"
	//也可以同时声明多个常量，如果省略了值则表示和上面一行的值相同
	const (
		c1 = "name"
		c2
		c3
		//c3 == c2 == c1
	)

	//iota是go语言的常量计数器，只能在常量的表达式中使用
	//iota在const关键字出现时将被重置为0，const中每新增一行常量声明，注意是每新增一行
	//将使iota计数一次(iota可理解为const语句块总的行索引)
	//也能使用 _ 跳过某些值
	const (
		counter1 = iota
		counter2 // = 1
		counter3 // = 2
		_
		counter4 // = 4
	)
	const (
		a1, a2 = iota + 1, iota + 2 //这时，a1 a2分别为1 ，2，因为在同一行，所以iota相同
	)
}

func opString() {

}

func main() {
	s := "Hello沙河小王子"
	c := 0
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			c++
		}
	}
	fmt.Print(c)
}

package main

import (
	"fmt"
)

//函数传递的都是值，也就是函数参数都是值传递，而不是引用传递，这点与java不同，java对于基本类型是值传递，而引用类型则是引用传递
//返回值可以命名，也可以不命名
//命名返回值就相当于在函数中声明一个变量
func sum(a int, b int) (ret int) {
	ret = a + b
	return
}

func sum1(a int, b int) int {
	return a + b
}

//也可以有多个返回值
func f1() (int, int) {
	return 0, 0
}

//参数类型简写。当参数中连续多个参数的类型一致时，我们可以讲非最后一个参数的类型省略
func f2(a, b int, c, d string, e, f bool) int {
	return a + b
}

//可变长参数
//可变长参数必须要放在函数参数最后，可变长参数其实就是一个切片
func f3(x int, y ...int) {

}

//defer语句
//defer语句会将起后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，讲延迟处理的语句按defer定义的逆序进行执行，也就是说
//先被defer的语句最后被执行，最后被defer的语句最先被执行
//defer多用于函数结束之前释放资源，(文件句柄，数据库连接，socket连接)

/**
在go中，reture语句不是原子操作，分为 给返回值赋值，和ret指令。
这defer语句的执行时机则是在这两步之间：
无defer语句时的执行步骤： return x 其实是 返回值 = x -> ret指令
有defer语句时的执行步骤： return x 其实是 返回值 = x -> 运行defer语句 -> ret指令
*/
func defFunc() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
	//该函数执行结果为
	/**
	start
	end
	3
	2
	1
	*/
}

func isPalindrome(s string) bool {
	for index := range s {
		if s[index] != s[len(s)-index-1] {
			return false
		}
	}
	return true
}
func main() {
	var s string = "how do you do"
	for i, c := range s {
		fmt.Println(i, c)
		fmt.Printf("%T %c\n", i, c)
	}
}

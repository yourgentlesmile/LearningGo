package main

import (
	"fmt"
)

//帮助手册：studygolang.com/pkgdoc

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

//分支选择语句
func switchFunc() {
	// fallthrough 表示命中了c这个分支后还能接着执行他下面的分支(d分支)
	// 功能就类似于java中的case里面不写break，会继续向下执行直到遇见break

	s := 'c'
	switch s {
	case 'c':
		fmt.Print("get c")
		fallthrough
	case 'd':
		fmt.Print("get d")
	}
}

// 循环语句
func loopFunc() {
	// 循环中断可以使用break，要是跳出多重for循环，可以使用goto标签
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(i + j)
			if j == 10 {
				goto out
			}
		}
	}
out:
	fmt.Print("I'm out.")
}

//运算符
func opFunc() {
	//基本运算符与其他语言一致，不再举例
	//不同点：自增语句不能跟在等号后
	//即 b = c++(或c--)是不允许的，这点与其他语言有所差别
	//只能单独存在，例如： c++
}

//数组
func arrayFunc() {
	//长度为3的整型数组
	var a [3]int
	fmt.Print(a[1])

	//有一点要注意，长度不同的数组变量不能互相赋值
	var b [4]int
	fmt.Print(b[1])
	//a = b 是不允许的，这是与java不同的地方
	//并且a与b是不同的类型
	//如果越界访问，会出现如下错误：000panic: runtime error: index out of range [3] with length 3

	//初始化 1
	var c [2]int = [2]int{1, 2}
	fmt.Print(c[1])
	//初始化 2，初始化数组长度可以由后面大括号的元素个数决定
	var d = [...]int{1, 2, 3, 4, 5}
	fmt.Print(d[1])

	//也可以部分初始化，未初始化到的则是默认值
	var e = [3]int{1, 2}
	fmt.Print(e[1])

	//也可以根据索引初始化
	//下标为0的初始化为1，下标为4的初始化为2
	var f = [5]int{0: 1, 4: 2}
	fmt.Print(f[1])

	//数组遍历
	for i := 0; i < len(f); i++ {
		fmt.Print(f[i])
	}
	for index, value := range f {
		fmt.Print(index, value)
	}

	//多维数组
	var all = [3][2]int{
		[2]int{1, 2},
		[2]int{1, 2},
		[2]int{1, 2}}
	var city = [3][2]string{
		{"", ""},
		{"", ""},
		{"", ""}}
	fmt.Print(all, city)
	//多维数组遍历
	for _, v1 := range city {
		fmt.Print(v1)
		for _, v2 := range v1 {
			fmt.Print(v2)
		}
	}

	//注意，数组是值类型而不是引用类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Print(b1, b2)
	//得到的是{1,2,3}, {100,2,3}
	//这就表明，数组是一个值类型；这和java完全不同，java中，数组是一个引用类型

	//同类型的数组支持== !=进行比较，会比较内容
	var q = [3][2]int{
		[2]int{1, 2},
		[2]int{1, 2},
		[2]int{1, 2}}
	var w = [3][2]int{
		[2]int{1, 3},
		[2]int{1, 2},
		[2]int{1, 2}}
	fmt.Print(q == w)

	//[n]*T 表示指针数组，*[n]T表示数组指针，这点和C++是一致
}

/*
切片是一个拥有相同类型元素的可变长度的序列，他是基于数组类型做的一层封装
它非常灵活，支持自动扩容。
注意：切片是一个引用类型，它的内部结构包含地址，长度和容量。切片一般用于快速地操作一块数据集合

*/
func sliceFunc() {

}

/**
指针与C语言的使用保持了一直，&为取地址符，*为地址取值
*/
func pointer() {
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) //输出*int表示int指针

	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m) //输出int表示int类型
}

/**
可以给指针一个已经初始化过的内存区域，用法也和C类似

new函数签名为func new(Type) *Type,表示接受一个类型，返回一个指向该类型内存的指针

make也是用于内存分配的，区别于new，它只用于slice，map，会议及chan(通道)的内存创建，而且它返回的类型就是这三个类型本身
而不是他们的指针类型，所以就没必要返回他们的指针了，make函数的签名如下：
func make(t Type, size ...IntegerType) Type
make函数是无可替代的，我们在使用slice，mao，以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作


区别:
1、二者都是用来做内存分配的。
2、make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3、而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
func makeAndNew() {
	var a = new(int) //new函数会申请一块内存
	*a = 100
	fmt.Println(*a)
	var b *int //此时， b = nil，是不能直接使用的
	fmt.Println(b)

	var c map[string]int
	c = make(map[string]int, 10)
	c["沙河娜扎"] = 100
	fmt.Println(c)
}

/**
hash表的使用
*/
func mapFunc() {
	var a map[string]int
	a = make(map[string]int, 10) //可以自动扩容
	a["a"] = 1
	a["b"] = 2
	fmt.Println(a["123"]) //如果不存在，则会返回对应类型的零值
	value, ok := a["13"]
	if !ok {
		fmt.Println("key不存在")
	} else {
		fmt.Println(value)
	}

	//map的遍历
	//遍历kv
	for k, v := range a {
		fmt.Println(k, v)
	}
	//遍历key
	for k := range a {
		fmt.Println(k)
	}
	//遍历value
	for _, v := range a {
		fmt.Println(v)
	}

	//删除map中的键值对
	delete(a, "a") //删除不存在的也不会报错

	//map和slice组合

	//切片中存map
	var s1 = make([]map[int]string, 0, 10) //第一个参数是切片的长度,第二个是底层数组的容量
	s1[0] = make(map[int]string, 10)
	//map中存切片
	var s2 = make(map[string][]int, 10)
	s2["shanghai"] = []int{1, 2, 3}
}
func main() {
	makeAndNew()
}

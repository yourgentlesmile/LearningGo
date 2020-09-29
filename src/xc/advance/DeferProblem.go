package main

//defer 延迟执行关键字，被修饰的语句会在返回值赋值之后，在方法返回之前执行

//Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
//在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
//也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

//5
func f1() int {
	x := 5
	defer func() { //匿名函数
		x++
	}()
	return x
}

//6
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

//5
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//5
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//func main() {
//	fmt.Println(f1())
//	fmt.Println(f2())
//	fmt.Println(f3())
//	fmt.Println(f4())
//}

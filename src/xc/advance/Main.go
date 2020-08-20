package main

import "fmt"

//使用append可以给切片添加元素，当容量不足时，会创建一个新的数组用于保存，将旧值搬到新的数组中
//append每次添加后会返回新的引用，这个引用推荐用原来的变量名接受，以免混乱
//关于slice的扩容策略，可以看slice.go的growslice函数

//切片总没有删除切片元素的专用方法，可以使用切片本身的特性去完成
func appendFunc() {
	a := []int{1, 2, 3, 4, 5}
	b := append(a, 4)
	fmt.Print(b)
	//删除下标为2的元素; 切片后面的三个点...表示参数张开，这个和js的参数展开是一样的 scala中的则是 :_*
	a = append(a[:2], a[3:]...)
	fmt.Print(a)
}

//copy(src ,dist) ，将src切片中的元素复制到切片c
func copyFunc() {
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 2, 2)
	copy(a, c)
}
func main() {
	appendFunc()
}

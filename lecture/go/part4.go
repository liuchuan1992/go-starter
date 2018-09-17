package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	for i := 0 ;i < len(arr);i++{
		fmt.Println(arr[i])
	}
	m1()
	m2()
}

func m1() {
	arr := [10]int{5:1}
	fmt.Println(arr)
}

/**
指向数组的指针和指针数组
 */
func m2()  {
	arr := [10]int{1:1}
	p := &arr //&取对象的指针地址
	fmt.Println(p)
	a := new([10]int)
	fmt.Println(a)
}
package main

import (
	"fmt"
	"strconv"
)

type (
	文本 string
)


func main() {
	//var i [9]int
	//fmt.Println(math.MaxInt8)
	//fmt.Println(i)
	//var str 文本
	//fmt.Print(str)
	//m1()
	//m2()
	//m3()
	m4()
	m5()
}

func m1()  {
	//变量声明
	var a int
	//变量赋值
	a = 123
	fmt.Println(a)
	//变量的声明同时赋值  全局变量不能用:=的方式 所以这个赋值一般是用在全局变量中
	var b int = 3
	//上行的格式可省略变量类型 由编译器推断
	var c = 123
	//变量声明与赋值的最简写法
	d := 123
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func m2() {
	var a,b,c,d = 1,2,3,4
	e,f,g,h := 5,6,7,8
	fmt.Println(a,b,c,d,e,f,g,h)
}

func m3() {
	var b int
	var a float32 = 1.2
	b = int(a)
	fmt.Println(b)
}

func m4() {
	var a int = 65
	b := string(a)
	fmt.Println(b)
}

/**
strconv这个函数方法
 */
func m5() {
	var a int = 65
	b := strconv.Itoa(a)
	fmt.Println(b)
}
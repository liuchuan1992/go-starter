package main

import "fmt"

const(
	a = 1
	e = 2
	//前面有两个变量 所以b从2开始了
	b = iota
	c
	d
)

/*
请结合常量与iota与<<运算符实现计算机储存单位的枚举
 */
const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	fmt.Println(a,b,c,d)
	fmt.Println("--------------")
	m1()
	m2()
	fmt.Println("--------------")
	m3()
	m4()
	fmt.Println("--------------m5")
	m5()
}

/**  位运算符
	6 :0110
	11:1011

& (和):两个是1取1 其余取0   0010  2
| (或）:如果有一个是1 就是1 1111 15
^ :两位只有一个是1时才成立   1101 13
&^ : 如果第二位是1时，把第一位的数字强制转为0 0101 4
 */
func m1() {
	fmt.Println( 6 & 11)
	fmt.Println( 6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println( 6 &^ 11)
}
func m2() {
	fmt.Println(B,KB,MB,GB)
}

func m3() {
	a := 1
	//定义了int类型的指针  &a 取变量的指针
	var p *int = &a
	fmt.Println(p)
	//*p取指针对应的变量的值
	fmt.Println(*p)
}

func m4() {
	for i := 0 ;i < 10;i++{
		fmt.Println(i)
	}
}

func m5() {
	a := 1
	str := "123"
	b := len(str)
	fmt.Println(b)
	for{
		a++
		if a > 10 {
			break
		}
	}
	fmt.Println(a)
}

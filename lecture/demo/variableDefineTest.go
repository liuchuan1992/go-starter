package main

import "fmt"

type S struct {
	s *A
}

type A struct {

}
/**
	变量的定义 包括普通的变量，数组 slice,map,等等
 */
func main() {
	testPoint()
}

func normalVar() {
	//先定义后赋值
	var i int
	i = 10;
	fmt.Println(i)
	//定义并赋值
	var j int = 10
	fmt.Println(j)
	//精简
	k := 10
	fmt.Println(k)
}

func arryVar() {
	//...会根据{}内元素的个数确定数组的长度
	arr := [...]int{}
	fmt.Println(arr)
	//指定长度
	a := new([2]int)
	fmt.Println(a)
}

func sliceVar(){
	//直接定义
	sli := []int{1,2,3}
	fmt.Println(sli)
	//数组截取
	a := [10]int{}
	sli1 := a[1:5]
	fmt.Println(sli1)
	//make定义
	sli2 := make([]int,5,10)
	fmt.Println(sli2)
}

func testPoint(){
	var i *int
	a := 10
	i = &a

	fmt.Println(i,*i)
}
func test()  {
	var i *int
	a := 10
	i = &a
	fmt.Println(*i)
}
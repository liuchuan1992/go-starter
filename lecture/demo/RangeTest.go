package main

import "fmt"

type Student struct {
	name string
	age int
	sex string
}
func main() {
	arr := [...]int{1,2,3,4,5}
	sum := 0
	for i,num := range arr{
		fmt.Println(i)
		sum += num
	}
	fmt.Println(sum)

	s := Student{"lvcf",27,"man	"}
	fmt.Println(s)
	s.plusAge()
	fmt.Println(s)
}

func init()  {
	fmt.Println("hello world init")
}

/**
加不加* 会影响到操作
	go语言是值传递 如果传过来是一个引用类型 那只是对引用类型的一个拷贝 方法内的修改不会影响到老对象
 */
func (s *Student) plusAge()  {
	s.age += 1
}
package main

import "fmt"

type _student struct {
	name string
	age int
	school string
}
func main()  {
	s1 := "浙江省"
	s2 := "杭州市"
	s3,s4 :=swap(s1,s2)
	fmt.Println(s3,s4)

	testStruct()

	student := _student{"lvcf",27,"zjut"}
	student.sayHello("lvcf")
}
/**
	函数多返回值
 */
func swap(s1,s2 string)(string,string)  {
	tmp := s1
	s1 = s2
	s2 = tmp
	return s1, s2
}

func testStruct(){
	s1 := _student{"lvcf",27,"zjut"}
	s1.name = "liuchuan"
	fmt.Println(s1)
}

func (student _student) sayHello(s string){
	fmt.Println("hello ",s)
}
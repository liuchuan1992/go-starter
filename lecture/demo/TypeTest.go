package main

import "fmt"

type STR string

func main()  {
	var s STR
	s = "123"
	fmt.Println(s)
	s.SayHello()
}


func (s STR) SayHello(){
	fmt.Println("hello")
}

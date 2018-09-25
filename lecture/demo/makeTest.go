package main

import "fmt"

func main()  {
	//使用make定义slice  make关键字只能用来给
	var sli []int = make([]int,0,10)
	fmt.Println(sli)

	arr := [10]int{}
	fmt.Println(arr)

}



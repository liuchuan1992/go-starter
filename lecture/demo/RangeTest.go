package main

import "fmt"

func main() {
	arr := [...]int{1,2,3,4,5}
	sum := 0
	for i,num := range arr{
		fmt.Println(i)
		sum += num
	}
	fmt.Println(sum)
}



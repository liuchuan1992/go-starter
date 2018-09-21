package main

import "fmt"

func main() {
	//createSlice()
	//createSliceByArray()
	//createSilceByMake()
	testSlice()
}

func createSlice()  {
	// []为slice [n]或者[...]为数组
	var s []int
	fmt.Println(s)
}

func createSliceByArray() {
	a := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(a)
	s := a[1:5] //包含起始索引不包括终止索引 1:5 相当于下标 [1,5)
	fmt.Println(s)
}

func createSilceByMake()  {
	s1 := make([]int,10,20)
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
}

func testSlice()  {
	s1 := [...]int{1,2,3,4,5,6,7,8,9}
	s2 := s1[1:5] //1,2,3,4
	s3 := s1[2:5]//2,3,4
	fmt.Println(s3)
	s2[1] = 100
	fmt.Println(s3)

}
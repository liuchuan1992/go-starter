package main

import "fmt"

func main()  {
	m1()
	m2()
	m3()
	m4()
	m5()
	m6()
}

/**
map创建
 */
func m1()  {
	var m map[int]string
	m = map[int]string{}
	fmt.Println(m)
	m1 := map[int]string{}
	fmt.Println(m1)
	m2 := make(map[int]string)
	fmt.Println(m2)
}

/**
map 添加元素
 */
func m2()  {
	m := make(map[int]string)
	m[0] = "lvcf"
	fmt.Println(m)
}

/**
	添加元素
 */
func m3(){
	m := map[string]string{}
	m["lili"] = "china"
	m["mike"] = "usa"
	fmt.Println(m)
}
/**
	通过key获取元素
 */
func m4() {
	m := map[string]string{}
	m["lili"] = "china"
	m["mike"] = "usa"
	fmt.Println(m["lili"])
}

func m5() {
	m := map[string]string{}
	m["lili"] = "china"
	m["mike"] = "usa"
	capital ,ok := m["minke"]
	if ok{
		fmt.Println(capital)
	}else{
		fmt.Println("首都不存在")
	}
}

func m6() {
	m := map[string]string{}
	m["lili"] = "china"
	m["mike"] = "usa"
	fmt.Println(m)
	delete(m,"lili")
	fmt.Println(m)
}
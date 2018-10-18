package main

import (
	"fmt"
	"strconv"
)

func main(){
	//	1.ParseBool 字符串转布尔
	// 返回一个boolean类型的值，接收以下字符串 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	//	其余的返回报错
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("t"))
	//	2.FormatBool 表达式转布尔
	// 根据传入的 表达式 返回 "true" or "false"
	fmt.Println(strconv.FormatBool( 1 >2))
	// 3.AppendBool
	result := make([]byte,0)
	result = strconv.AppendBool(result,1<2)
	fmt.Printf("%s\n", result) // true
	// 4.ParseFloat
	fmt.Println(strconv.ParseFloat("0.12345678901234567890",32))
	// 5.ParseInt 将制定进制的字符串转化为10进制
	fmt.Println(strconv.ParseInt("12345",8,16))

}

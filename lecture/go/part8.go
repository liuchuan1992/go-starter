package main

import "fmt"

/**
	定义一个接口 USB 包含: Name和Conncet方法
 */
type USB interface {
	Name() string
	Connector
}
type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func main() {
	var u USB
	u = PhoneConnector{"phoneConnector"}
	u.Connect()
	fmt.Println(u.Name())

	Disconnect(u)


	method()
}

func (pc PhoneConnector) Name() string  {
	return pc.name
}

func (pc PhoneConnector)Connect()  {
	fmt.Println(pc.name,"connected success.")
}

/**
	Comma-ok断言的语法是：value, ok := element.(T)。
	element必须是接口类型的变量，T是普通类型。
	如果断言失败，ok为false，否则ok为true并且value为变量的值。
 */
func Disconnect(usb USB) {
	if pc,ok := usb.(PhoneConnector); ok{
		fmt.Println("Disconnected:",pc.name)
		return
	}
	fmt.Println("未知设备",usb)
}

/**
	switch type
 */
func DisConnect(usb interface{})  {
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected:",v.name)
	default:
		fmt.Println("未知设备",usb)
	}
}

func method() {
	pc := PhoneConnector{"phoneConnector"}
	var a Connector
	a = Connector(pc)
	a.Connect()
}
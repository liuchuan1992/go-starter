#go基本类型
1.布尔型 bool
取值范围 true,false
注意事项 不能用0 1代表

2.整型 int /uint
根据平台可能为64位或32位

3.8位整型 int8 /uint8
长度:1个字节 或者 8位

4.byte类型  字节类型
(uint8的别名)

5.16位整型 int16/uint16
长度:2个字节

6.32位整型 int32/uint32
  长度:4个字节
(rune 是 int32的别名)

7.64位整型 int64/uint64
长度:8字节

8.浮点型
float32 /float64
长度4/8字节
小数位：精确到7/15位

double 类型是没有的

9.其他值类型
array struct string

10.引用类型
slice map chan

11.接口类型 interface

12.函数类型 func

#类型零值

math.maxInt8 可以拿到int8的最大值

#类型别名
type(
  byte uint8
  rune int32
  文本 string
)

#变量的声明与赋值
1.变量的声明格式
var <变量名称> <变量类型>
*  var a int //变量的声明

2.变量的赋值格式
* <变量名称> = <表达式>
a  = 123

3.声明的同时赋值
* var <变量名称> [变量类型] = <表达式>
var b int = 123

4.多个变量的声明与赋值
全局变量的声明可以用var()的方式进行简写，但是局部变量不行
全局变量的声明不能省略var，但可以并行方式 所谓的并行方式是指 a,b = 1,2
所有的变量都可以使用类型推断
局部变量不可以使用var()方式简写，只能使用并行方式
其实 := 中的冒号是用来替代var关键字的
空白符号 _


#go语言类型转换
go不存在隐式转换 所有类型必须是显式声明
转换只能发生在两种相互兼容的类型之间
格式：
<ValueA> [:] = <TypeOfValueA>(<VaueB>)
例1:
var a float32 = 1.1
b := int(a)
例2:
var b int 
var a = float32 =1.1
b = int(a)

上面两个例子中何时可以去掉: 如果ValueA之前没有赋过值 那就需要:  :相当于var

注意的是 bool和整型之间是不能进行类型转换
 













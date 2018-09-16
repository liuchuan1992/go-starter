#part1
1.本地打开go的api文档
godoc -http=:6060

2.go 内置关键字

3.go程序结构
go程序是通过package

4.编译优化
没有用到的包编译的时候会提示必须移除

5.main包main方法是程序的入口

6.package别名 
import std "fmt"
std.Print(...)
import . "fmt"
Print(...) 省略调用  不太建议

7.可见性规则
go语言中根据大小写来决定该常量，变量，类型，接口，结构或函数否可以被外部包调用
根据约定，已大写字母开头表示public，小写字母开头表示private

总结
package -> import -> const -> var -> type -> interface -> function ...

package main

/**
mysql 测试连接
 */
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)




func main(){
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	//关闭数据库，db会被多个goroutine共享，可以不调用

	res, err := stmt.Exec("Test", "people", "2017-10-27")
	id, err := res.LastInsertId()

	fmt.Println(id)
	defer db.Close()
}

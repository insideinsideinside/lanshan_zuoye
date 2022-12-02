package main

import (
	"database/sql" //标准库
	"fmt"
	_ "github.com/go-sql-driver/mysql" //我们使用的mysql，需要导入相应驱动包，否则会报错
	"log"
)

// 定义一个全局对象db
var db *sql.DB

func initDB() {
	var err error
	// 设置一下dns charset:编码方式 parseTime:是否解析time类型 loc:时区
	dsn := "root:wzy20040525@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开mysql驱动
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}

func main() {
	//初始化连接
	initDB()
	//var st Student
	//i := 0
	//for i = 0; i < 10; i++ {
	//	fmt.Printf("请输入id，名字和性别\n")
	//	fmt.Scanf("%d %s %s\n", &st.ID, &st.Name, &st.Gender)
	//	InsertStudent(st)
	//}
	queryMultiRowDemo()

}
func queryMultiRowDemo() {
	sqlStr := "select id, name, gender from student where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u Student
		err := rows.Scan(&u.ID, &u.Name, &u.Gender)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s gender:%s\n", u.ID, u.Name, u.Gender)
	}
}

type Student struct {
	ID     int64
	Name   string
	Gender string
}

func InsertStudent(st Student) {
	sqlStr := "insert into student(id,Name,gender) values (?,?,?)"
	_, err := db.Exec(sqlStr, st.ID, st.Name, st.Gender)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

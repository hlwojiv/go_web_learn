package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type myField struct {
	uid                           int
	username, created, department string
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=postgres sslmode=disable")
	checkErr(err)

	// 插入数据
	// var lastInsertId int
	// err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "hlwojiv", "搞事部门", "2019-12-09").Scan(&lastInsertId)
	// checkErr(err)
	// fmt.Println("最后插入的 id =", lastInsertId)

	// 更新数据
	// stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	// checkErr(err)

	// res, err := stmt.Exec("hlwojivv", 1)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// 查询数据
	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)

	// for rows.Next() {
	// 	data := new(myField)
	// 	err = rows.Scan(&data.uid, &data.username, &data.department, &data.created)
	// 	checkErr(err)
	// 	fmt.Printf("%v\t%v\t%v\t%v", data.uid, data.username, data.department, data.created)
	// 	fmt.Println()
	// }

	// 删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	attect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(attect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

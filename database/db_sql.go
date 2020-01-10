package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type myField struct {
	uid                           int
	username, created, department string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)

	// 插入数据
	// res := myInsert(db)
	// fmt.Println(res)

	// 更新数据
	// res := myUpdate(db)
	// fmt.Println(res)

	// 查询数据
	// mySelect(db)

	// 删除数据
	stmt, err := db.Prepare("DELETE from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(2)
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(id)

	db.Close()
}

func mySelect(db *sql.DB) {
	rows, err := db.Query("Select * from userinfo")
	checkErr(err)

	for rows.Next() {
		// var data myField
		data := new(myField)
		err = rows.Scan(&data.uid, &data.username, &data.department, &data.created)
		checkErr(err)
		fmt.Printf("%v\t%v\t%v\t%v", data.uid, data.username, data.department, data.created)
		fmt.Println()
	}
}

func myUpdate(db *sql.DB) int64 {
	stmt, err := db.Prepare("UPDATE userinfo set uid=? where username=?")
	checkErr(err)

	res, err := stmt.Exec(2, "dlnyoo")
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	return id
}

func myInsert(db *sql.DB) int64 {
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("dlnyoo", "扑街部门", "2019-12-24")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return id
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

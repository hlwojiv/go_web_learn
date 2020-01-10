package main

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Userinfo struct {
	Uid        int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}

type User struct {
	Uid     int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Tag))
}

func main() {
	o := orm.NewOrm()

	var user User

	user.Name = "dlnyoo"
	user.Departname = "搞事部"

	// user := User{Name: "hlwojiv"}

	// // 插入表
	// id, err := o.Insert(&user)
	// fmt.Printf("ID : %d, err: %v\n", id, err)

	// 更新表
	// user.Name = "firefox"
	// num, err := o.Update(&user)
	// fmt.Printf("NUM : %d, err: %v\n", num, err)

	// 读取 one
	// u := User{Id: 4}
	// err := o.Read(&u)
	// fmt.Printf("ERR: %v\n", err)

	// // 删除表
	// num, err := o.Delete(&u)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

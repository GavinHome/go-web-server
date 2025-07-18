package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct

type Userinfo struct {
	Uid        int64 `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Department string
	Created    time.Time
}

// type User struct {
// 	Uid     int `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
// 	Name    string
// 	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
// 	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
// }

// type Profile struct {
// 	Id   int
// 	Age  int16
// 	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
// }

// type Post struct {
// 	Id    int
// 	Title string
// 	User  *User  `orm:"rel(fk)"`
// 	Tags  []*Tag `orm:"rel(m2m)"` //设置一对多关系
// }

// type Tag struct {
// 	Id    int
// 	Name  string
// 	Posts []*Post `orm:"reverse(many)"`
// }

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8", 30)

	// 注册定义的 model
	orm.RegisterModel(new(Userinfo))
}

func main() {
	o := orm.NewOrm()
	var user Userinfo
	user.Username = "zxxx"
	user.Department = "zxxx"
	// 插入
	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}

	// 更新
	user = Userinfo{Uid: id}
	if o.Read(&user) == nil {
		user.Username = "MyName"
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
		}
	}

	// 查询
	user = Userinfo{Uid: id}
	err = o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Uid, user.Username)
	}

	// 删除
	if num, err := o.Delete(&Userinfo{Uid: id}); err == nil {
		fmt.Println(num)
	}
}

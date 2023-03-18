package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 关联一对一
// https://gorm.io/zh_CN/docs/has_one.html
var db *gorm.DB

func init() {
	var err error
	dns := "root:123@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func run() {
	// ===============创建数据=========
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&CreditCard{})
	//u1 := &User{
	//   Name: "swt",
	//   CreditCard: CreditCard{
	//      Number: "110",
	//   },
	//}
	//db.Create(u1)
	// ===============查询数据=========
	user := &User{}
	//user.ID = 2
	//var card CreditCard
	db.Preload("CreditCard").First(&user)
	fmt.Println(user.CreditCard.Number)
	//SELECT * FROM credit_cards WHERE user_id = 2;

}
func main() {
	run()
}

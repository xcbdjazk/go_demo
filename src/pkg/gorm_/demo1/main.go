package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 关联一对一
var db *gorm.DB

//https://gorm.io/zh_CN/docs/belongs_to.html

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
	CompanyID  int
	Company    Company
	Company2ID int
	Company2   Company2
}

type Company struct {
	ID   int
	Name string
}

type Company2 struct {
	ID   int
	Name string
}

func run() {
	//user := User{
	//   Name: "123",
	//
	//}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Company{})
	db.AutoMigrate(&Company2{})
	// BEGIN TRANSACTION;
	// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
	// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "languages" ("name") VALUES ('ZH'), ('EN') ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "user_languages" ("user_id","language_id") VALUES (111, 1), (111, 2) ON DUPLICATE KEY DO NOTHING;
	// COMMIT;
	var u User
	var u2 User

	// 预加载，会加载相关的关联数据
	db.Preload("Company").First(&u)
	fmt.Println("pp", u.Name)
	fmt.Println("pp", u.Company.Name)
	db.Joins("Company").First(&u2)
	fmt.Println("jj", u2.Name)
	fmt.Println("jj", u2.Company.Name)
	//com := &Company{Name: "123"}
	//db.Create(&com)
	//fmt.Print(com.ID)
	//user.Company = *com
	//db.Create(&user)
}
func main() {
	run()
}

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 关联多对多
// https://gorm.io/zh_CN/docs/has_many.html

var db *gorm.DB

func init() {
	var err error
	dns := "root:123@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type Person struct {
	ID        uint
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses"`
}

type Address struct {
	ID   uint
	Name string
}

type PersonAddress struct {
	PersonID  uint
	AddressID uint
	ID        uint `gorm:"unique;autoIncrement:true"`
}

func run() {
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&PersonAddress{})
	db.AutoMigrate(&Address{})
	p := Person{Name: "123", Addresses: []Address{{Name: "1231"}, {Name: "1232"}}}
	db.Save(&p)

}
func main() {
	run()
}

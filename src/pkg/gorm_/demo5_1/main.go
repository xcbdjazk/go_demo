package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

type PersonAddress struct {
	PersonID  int
	AddressID int
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//func (PersonAddress) BeforeCreate(db *gorm.DB) error {
//	return db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})
//
//}
func run() {
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&PersonAddress{})

}
func main() {
	run()
}

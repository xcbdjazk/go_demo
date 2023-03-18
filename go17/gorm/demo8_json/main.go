package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dns := "root:123@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// User 用户
type User struct {
	gorm.Model
	Name    string
	Profile Profile `gorm:"type:json;comment:'个人信息'"json:"profile"`
}

// Profile 个人信息
type Profile struct {
	Email   string `json:"email"`
	PhoneNo string `json:"phoneNo"`
}

// Value 实现方法
func (p Profile) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现方法
func (p *Profile) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), p)
}

func main() {
	db.AutoMigrate(&User{})
	u := User{
		Name: "maocat",
		Profile: Profile{
			Email:   "maocat@gmail.com",
			PhoneNo: "18888888888",
		},
	}
	db.Save(&u)

	u.Profile.PhoneNo = "13666666666"
	db.Save(&u)

	var u1 User
	db.
		Where("profile->'$.email' = (?)", "maocat@gmail.com").
		First(&u1)
	fmt.Println(u1.Name)
	fmt.Println(u1.Profile)
}

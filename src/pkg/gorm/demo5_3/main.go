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

type Tag struct {
	//gorm.Model
	ID     uint   `gorm:"primaryKey"`
	Locale string `gorm:"primaryKey"`
	Value  string
}

type Blog struct {
	ID         uint   `gorm:"primaryKey"`
	Locale     string `gorm:"primaryKey"`
	Subject    string
	Body       string
	Tags       []Tag `gorm:"many2many:blog_tags;"`
	LocaleTags []Tag `gorm:"many2many:locale_blog_tags;ForeignKey:id,locale;References:id"`
	SharedTags []Tag `gorm:"many2many:shared_blog_tags;ForeignKey:id;References:id"`
}

func run() {
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Blog{})

}
func main() {
	run()
}

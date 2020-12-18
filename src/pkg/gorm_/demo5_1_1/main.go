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

type Product struct {
	ProductID       int              `gorm:"column:product_id;primary_key" json:"product_id"`
	Name            string           `gorm:"column:name" json:"name"`
	Categories      []Category       `gorm:"many2many:product_category;foreignkey:product_id;association_foreignkey:category_id;association_jointable_foreignkey:category_id;jointable_foreignkey:product_id;"`
	AttributeValues []AttributeValue `gorm:"many2many:product_attribute;foreignkey:product_id;association_foreignkey:attribute_value_id;association_jointable_foreignkey:attribute_value_id;jointable_foreignkey:product_id;"`
}

type Category struct {
	CategoryID int       `gorm:"column:category_id;primary_key" json:"category_id"`
	Name       string    `gorm:"column:name" json:"name"`
	Products   []Product `gorm:"many2many:product_category;foreignkey:category_id;association_foreignkey:product_id;association_jointable_foreignkey:product_id;jointable_foreignkey:category_id;"`
}

type AttributeValue struct {
	AttributeValueID int       `gorm:"column:attribute_value_id;primary_key" json:"attribute_value_id"`
	AttributeID      int       `gorm:"column:attribute_id" json:"attribute_id"`
	Value            string    `gorm:"column:value" json:"value"`
	Products         []Product `gorm:"many2many:product_attribute;foreignkey:attribute_value_id;association_foreignkey:product_id;association_jointable_foreignkey:product_id;jointable_foreignkey:attribute_value_id;"`
}

func run() {
	//db.AutoMigrate(&PersonAddress{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&AttributeValue{})
	//db.AutoMigrate(&Address{})

}
func main() {
	run()
}

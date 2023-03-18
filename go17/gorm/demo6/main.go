package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// 关联一对多
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

type User struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

func run() {
	// ===============创建数据=========
	//db.AutoMigrate(&User{})
	//user := User{
	//	Name: "jinzhu",
	//	Languages: []*Language{
	//		{Name: "ZH"},
	//		{Name: "EN"},
	//	},
	//}
	//
	//db.Create(&user)
	//db.Create(u1)
	// ===============查询数据=========
	var user User
	//db.First(&user)
	//var languages []Language
	//err := db.Model(&user).Association("Languages").Error
	//fmt.Println(err)
	//
	////db.Model(&user).Association("Languages").Find(&languages)
	////codes := []string{"jinzhu", "EN"}
	//db.Model(&user).Association("Languages").Find(&languages)
	//
	//fmt.Println(languages)
	//fmt.Println(user)
	//db.Model(&user).Where(1).First(&user)
	//fmt.Println(user.Name)
	//_ = db.Debug().Model(&user).Where(map[string]interface{}{
	//	"Name":"1",
	//}).Where(map[string]interface{}{
	//	"id":"1",
	//}).Where(map[string]interface{}{
	//	"id between @a1 and @a2": map[string]interface{}{
	//	"a1":"0",
	//	"a2":"2",
	//	},
	//}).First(&user)
	//r, _ := db.Rows()

	//r.Next()
	locking := clause.Locking{Strength: "UPDATE"}

	tx := db.Begin()
	tx.Debug().Clauses(locking).First(&user, 1)
	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	go func() {
		//time.Sleep(time.Second * 2)
		var user1 User
		e := db.First(&user1, 2).Update("name", "1").Error
		if e != nil {
			fmt.Println(e)
		}
	}()
	// 遇到错误时回滚事务
	//tx.Rollback()
	time.Sleep(time.Second * 3)
	// 否则，提交事务
	tx.Commit()
	//_ = db.Debug().Model(&user).Select(" * ").
	//	Where("TO_DAYS(created_at) = TO_DAYS(NOW())").First(&user)
	//fmt.Println(user)

	//var user User
	//db.Preload("Languages").First(&user)
	//fmt.Println(user.Languages[0].Name)
}
func main() {
	run()
}

package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"opensvc/pkg/setting"
)

var Db *gorm.DB

type Model struct {
	ID       int `gorm:"primary_key; AUTO_INCREMENT; column:id"`
	Region	 string `gorm:"column:region"`
	Auth_url string	`gorm:"column:auth_url"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Domain   string `gorm:"column:domain"`
	Project  string `gorm:"column:project"`
	Deleted  int 	`gorm:"column:deleted"`

}



func Setup() {
	var err error
	Db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.URL,
			setting.DatabaseSetting.Dbname))
	if err != nil {
		log.Fatalf("model.Setup err : %v", err)
	}
	if Db.HasTable(Model{}) == false {
		Db.CreateTable(Model{})
	}
	if Db.HasTable(Vm{}) == false{
		Db.CreateTable(Vm{})
	}



}
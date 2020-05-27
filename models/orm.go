package models

import (
	"fmt"
	"log"

	"github.com/dzlzh/gin-example/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	username := utils.Config.Section("mysql").Key("username").String()
	password := utils.Config.Section("mysql").Key("password").String()
	dbHost := utils.Config.Section("mysql").Key("dbHost").String()
	dbName := utils.Config.Section("mysql").Key("dbName").String()
	dbPrefix := utils.Config.Section("mysql").Key("dbPrefix").String()

	DB, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		dbHost,
		dbName))

	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbPrefix + defaultTableName
	}
}

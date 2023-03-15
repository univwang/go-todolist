package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	db.LogMode(true)
	if gin.Mode() == gin.ReleaseMode {
		db.LogMode(false)
	}
	db.SingularTable(true)

	DB = db
	migration()
}

package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go_gin/pkg/setting"
	"log"
)

var db *gorm.DB

/**
基类
*/
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbname, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("fail to parse database %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbname = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbname))
	if err != nil {
		log.Println(err)
		return
	}



	// 返回表名称
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	db.Close()
}

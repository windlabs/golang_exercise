package common

import (
	"fmt"
	"net"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	globalDb *gorm.DB
)

func GetDb()(*gorm.DB, error){
	if globalDb == nil {
		fmt.Println("init mysql connect")
		db, err := MysqlInit()
		globalDb = db
		return db, err
	}
	return globalDb, nil
}

func MysqlInit() (*gorm.DB, error){
	db,err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True",
		"go_user","go_user",net.JoinHostPort("127.0.0.1","3306"),"go_db"))
	if err != nil {
		fmt.Println("db init error", err, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=local",
			"go_user","go_user",net.JoinHostPort("127.0.0.1","3306"),"go_db"))
		return nil, err
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetConnMaxLifetime(time.Second * 10)
	return db, err
}
package orm

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "rpa:rpa@rpa.com@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("链接数据库失败")
	}
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatal("查表错误")
	}
	for _, table := range tables {
		fmt.Println(table)
	}
	fmt.Println("-------------------")
	var tables2 []string
	db.Raw("show tables").Scan(&tables2)
	for _, table := range tables2 {
		fmt.Println(table)
	}
}

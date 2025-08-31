package raw

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"example/db/common"
)

// 获取数据库连接
func GetDbConn(dbType string) *sql.DB {
	dsn := common.GetDbDsn(dbType)
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		log.Fatal("连接数据库失败：", err)
	}
	return db
}

// 创建数据库表
func CreatePostsTable(db *sql.DB) {
	createSQL := `
	create table if not exists posts (
		id int not null auto_increment,
		title varchar(255) comment "标题",
		content varchar(255) comment "内容",
		primary key(id)
	)
	`
	_, err := db.Exec(createSQL)
	if err != nil {
		log.Fatal("初始化数据表posts失败", err)
	}
	fmt.Println("初始化posts数据表成功")
}

// 插入一条贴文数据
func InsertAPost(db *sql.DB, id uint, title string, content string) {
	// 先判断贴文是否存在
	exists := QueryPostById(db, id)
	if exists {
		fmt.Println("该ID的贴文已经存在")
		return
	}
	// 执行插入逻辑
	insertSQL := `
	insert into posts (id, title, content)
	values(?, ?, ?)
	`
	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatal("准备SQL时出现错误：", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(id, title, content)
	if err != nil {
		log.Fatal("执行SQL出现错误", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("获取执行结果失败：", err)
	}
	fmt.Println("执行成功，受影响行数：", rowsAffected)
}

// 根据ID查询贴文
func QueryPostById(db *sql.DB, id uint) bool {
	var exists bool
	querySQL := "select exists(select 1 from posts where id = ?)"
	err := db.QueryRow(querySQL, id).Scan(&exists)
	if err != nil {
		log.Fatal("查询贴文是否存在时出现错误：", err)
	}
	return exists
}

// 查询所有的贴文数据
func QueryAllPost(db *sql.DB) {
	querySQL := "select id, title, content from posts"
	result, err := db.Query(querySQL)
	if err != nil {
		log.Fatal("执行查询贴文时失败：", err)
	}
	for result.Next() {
		var id uint
		var title string
		var content string
		err := result.Scan(&id, &title, &content)
		if err != nil {
			log.Fatal("扫描查询结果时出现错误：", err)
		}
		fmt.Printf("ID=%d\t标题=%s\t内容=%s\n", id, title, content)
	}
}

// 根据ID删除贴文
func DeletePostById(db *sql.DB, id uint) {
	if !QueryPostById(db, id) {
		fmt.Printf("ID为%d的贴文不存在。", id)
	}
	deleteSQL := "delete from posts where id = ?"
	_, err := db.Exec(deleteSQL, id)
	if err != nil {
		fmt.Println("删除贴文时失败：", err)
		return
	}
	fmt.Printf("删除贴文[ID=%d]失败：", id)
}

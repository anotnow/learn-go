package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func demo_sqlite_conn() *sql.DB {
	sqliteDbPath := "./db/sqlite-demo.db"
	db, err := sql.Open("sqlite3", sqliteDbPath)
	if err != nil {
		log.Fatal("打开sqlite数据库失败", err)
	}
	fmt.Println("打开sqlite数据库成功")

	// 测试数据库连通性
	err = db.Ping()
	if err != nil {
		log.Fatal("无法连接到数据库: ", err)
	}
	fmt.Println("连接数据库成功")

	return db
}

func demo_sqlite_create() {
	db := demo_sqlite_conn()
	defer db.Close()
	initSQL := "create table posts(id int, title text, content text)"
	stmt, err := db.Prepare(initSQL)
	if err != nil {
		log.Fatal("初始化sql语句失败", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		log.Fatal("执行sql失败")
	}
	fmt.Println("sql执行结果: ", result)

}

func demo_sqlite_insert() {
	db := demo_sqlite_conn()
	defer db.Close()

	insertSQL := "insert into posts(id, title, content) values(?, ?, ?)"
	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatal("准备sql失败：", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(1, "中国必胜", "钓鱼岛是中国的")
	if err != nil {
		log.Fatal("插入失败：", err)
	} else {
		fmt.Println("数据插入成功")
	}
	_, err = stmt.Exec(2, "铭记历史", "勿忘国耻，东方必将崛起")
	if err != nil {
		log.Fatal("插入失败：", err)
	} else {
		fmt.Println("数据插入成功")
	}
}

func demo_sqlite_query() {
	db := demo_sqlite_conn()
	defer db.Close()

	querySQL := "select id, title, content from posts"
	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatal("查询数据失败", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var content string
		err := rows.Scan(&id, &title, &content)
		if err != nil {
			log.Fatal("扫描查询结果失败：", err)
		}
		fmt.Printf("ID=%d, 标题=%s, 内容=%s\n", id, title, content)
	}
}

func demo_sqlite_update() {
	db := demo_sqlite_conn()
	defer db.Close()

	updateSQL := "update posts set id = ? where id = ?"
	stmt, err := db.Prepare(updateSQL)
	if err != nil {
		log.Fatal("准备sql失败：", err)
	}
	result, err := stmt.Exec(2, 1)
	if err != nil {
		log.Fatal("执行sql失败：", err)
	} else {
		fmt.Println("执行sql成功, ", result.RowsAffected)
	}
}

func main() {
	// db := demo_sqlite_conn()
	// db.Close()
	// demo_sqlite_create()
	// demo_sqlite_insert()
	demo_sqlite_update()
	demo_sqlite_query()
}

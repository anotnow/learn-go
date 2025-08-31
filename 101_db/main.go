package main

import (
	"example/db/raw"
	"fmt"
	"os"
	"strconv"
)

func DemoDb() {
	db := raw.GetDbConn("mysql")
	raw.CreatePostsTable(db)
	var choice string
	for {
		showMenu()
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			raw.QueryAllPost(db)
		case "2":
			fmt.Print("请输入要查询的ID")
			var targetId string
			fmt.Scanln(&targetId)
			convID, _ := strconv.ParseUint(targetId, 10, 32)
			raw.DeletePostById(db, uint(convID))
		case "3":
			fmt.Println("你输入了3")
		case "4":
			fmt.Println("你输入了4")
		case "q":
			fmt.Println("正在退出")
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}
	}
}

func showMenu() {
	fmt.Println()
	fmt.Println("------功能菜单------")
	fmt.Println("1. 查询所有贴文")
	fmt.Println("2. 根据ID查询贴文")
	fmt.Println("3. 根据ID删除贴文")
	fmt.Println("4. 插入一条新的贴文")
	fmt.Println("q. 退出")
	fmt.Println("------------------")
	fmt.Print("请输入你的选择：")
}

func main() {
	DemoDb()
}

package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// 最简用法
func SimpleDemo1() {
	// 设置运行模式（可选）
	gin.SetMode(gin.ReleaseMode)
	// 实例化路由对象
	router := gin.Default()
	// 绑定路径和处理函数
	router.GET("/foo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"foo": "bar",
		})
	})
	// 运行
	router.Run("127.0.0.1:8888")
}

// 使用http包的响应码常量
func SimpleDemo2() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/foo", func(c *gin.Context) {
		// 此处使用http响应码常量
		c.JSON(http.StatusOK, gin.H{
			"foo": "bar",
		})
	})
	router.Run("localhost:9999")
}

// 路由分组
func RouterGroupDemo() {
	SimpleEcho := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"foo": "bar",
		})
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1 := router.Group("v1")
	v1.GET("page1", SimpleEcho) // /v1/page1
	v1.GET("page2", SimpleEcho) // v1/page2

	v2 := router.Group("v2")
	v2.GET("page1", SimpleEcho) // /v2/page1
	v2.GET("page2", SimpleEcho) // /v2/page2

	router.Run("127.0.0.1:8001")
}

func main() {
	var wg sync.WaitGroup
	// 正常情况下会Add指定的协程数量
	// 然后协程函数结束时调用wg.Done
	// 此处添加了1个等待
	// 但是没有任何函数调用wg.Done
	// 所以会一直阻塞
	wg.Add(1)
	go SimpleDemo1()
	go SimpleDemo2()
	go RouterGroupDemo()
	// 声明阻塞
	wg.Wait()
}

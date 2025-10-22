# 创建并运行 GO 项目

## 1. 初始化 Go 包

执行如下命令可初始化一个Go包

```cmd
go mod init example.com/xxx/yyy
```

上面命令会生成 `go.mod` 文件，内容示例如下：

```go
module example.com/xxx/yyy

go 1.24.5
```


## 2. 创建 Go 程序

1. 创建一个Go文件 `xxx.go`（名字任意），并写入以下内容。
2. 在包内执行 `go run .` 或者 `go run xxx.go` 即可执行程序。

```go
package main

import "fmt"

func main() {
	fmt.Println("hello,world")
}
```
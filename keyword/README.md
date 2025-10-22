# Go 语言的语法

参考文档：[The Go Programming Language Specification](https://go.dev/ref/spec#Types)

## Go 语言的关键字(25个)

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## Go 语言的保留字(37个)

```
# 常量
true、false、iota、nil

# 类型
int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、uintptr、
float32、float64、complex64、complex128、bool、byte、rune、string、error

# 函数
make、len、cap、new、append、copy、close、delete、complex、real、imag、panic、recover
```

## 关键字和保留字的区别

维度 | 关键字（Keywords） | 保留字（Reserved Words）
-|-|-
当前作用 | 直接参与语法逻辑，有明确含义 | 无实际语法功能，仅被预留
使用限制 | 不可作为标识符 | 不可作为标识符
目的 | 构建语言基本语法结构 | 为未来语法扩展预留，避免命名冲突

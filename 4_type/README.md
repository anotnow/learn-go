# Go的数据类型

参考文档：[The Go Programming Language Specification](https://go.dev/ref/spec#Types)

类型 | 长度(字节) | 默认值 | 说明
-|-|-|-
bool | 1 | false | 
byte | 1 | 0 | uint8
rune | 4 | 0 | Unicode Code Point, int32
int, uint | 4或8 | 0 | 32 或 64 位
int8, uint8 | 1 | 0 | -128 ~ 127, 0 ~ 255，byte是uint8 的别名
int16, uint16 | 2 | 0 | -32768 ~ 32767, 0 ~ 65535
int32, uint32 | 4 | 0 | -21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名
int64, uint64 | 8 | 0 | 
float32 | 4 | 0.0 | 
float64 | 8 | 0.0 | 
complex64 | 8 |  | 
complex128 | 16 |  | 
uintptr | 4或8 |  | 存储指针的 uint32 或 uint64 整数
array |  |  | 值类型
struct |  |  | 值类型
string |  | "" | UTF-8 字符串
slice |  | nil | 引用类型
map |  | nil | 引用类型
channel |  | nil | 引用类型
interface |  | nil | 接口
function |  | nil | 函数

## Boolean 布尔型

`true` 和 `false`

## Numeric 数值型

整数、浮点数、复数。

```go
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE 754 32-bit floating-point numbers
float64     the set of all IEEE 754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

```go
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

## String 字符串型

一个字符串值代表一个字节序列，字节的的个数称为字符串的长度（永远非负）。
字符串中的字节可以通过整数索引（`0~len(s)-1`）访问，但是无法访问到这个字节的内存地址（ `&s[i]` 会报错）。

## Array 数组型

一个数组是一个单一元素类型的序列，元素的个数成为元素的长度。
长度是数字类型的一部分，数组类型总是一维的，但是可以组合成多维。

```go
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

## Slice 切片型

切片可以认为是可变数组。

## Struct 结构体

结构体是一系列元素的集合。

```go
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

## Pointer 指针型

## Function 函数型

## Interface 接口型

## Map 映射型

## Channel 通道型



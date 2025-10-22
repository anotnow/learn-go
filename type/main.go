package main

import (
	"fmt"
	"math"
	"unsafe"
)

func demo_int() {
	fmt.Println("=== 使用 math 包 ===")
	fmt.Printf("math.MaxInt: %d\n", math.MaxInt)
	// fmt.Printf("math.MaxUint: %v\n", math.MaxUint)

	fmt.Println("\n=== 平台相关信息 ===")
	fmt.Printf("int 大小: %d 字节\n", unsafe.Sizeof(int(0)))
	fmt.Printf("uint 大小: %d 字节\n", unsafe.Sizeof(uint(0)))

	fmt.Println("\n=== 所有整数类型的范围 ===")
	fmt.Printf("int8:  %d 到 %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16: %d 到 %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32: %d 到 %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64: %d 到 %d\n", math.MinInt64, math.MaxInt64)

	fmt.Printf("uint8:  0 到 %d\n", math.MaxUint8)
	fmt.Printf("uint16: 0 到 %d\n", math.MaxUint16)
	fmt.Printf("uint32: 0 到 %d\n", math.MaxUint32)
	// fmt.Printf("uint64: 0 到 %d\n", math.MaxUint64)
}

func main() {
	demo_int()
}

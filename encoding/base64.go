package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

// 字符串转base64
func ToString() {
	var foo = "你好bar1234"
	var result string = base64.StdEncoding.EncodeToString([]byte(foo))
	fmt.Println(result)
}

// base64转字符串
func FromString() {
	var base64Str string = "5L2g5aW9YmFyMTIzNA=="
	result, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
}

func Base64Demo() {
	ToString()
	FromString()
}

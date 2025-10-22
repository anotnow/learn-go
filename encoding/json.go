package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  uint8
}

var person Person
var resultFile string

func init() {
	person.Name = "dev"
	person.Age = 18
	resultFile = "./result.json"
}

// 结构体 -> 文件
func Dump() {
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(file).Encode(person)
	if err != nil {
		log.Fatal(err)
	}
}

// 结构体 -> JSON字符串
func Dumps() {
	result, err := json.MarshalIndent(person, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
}

// 文件 -> 结构体
func Load() {
	file, err := os.Open(resultFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(file).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
}

// 字符串 -> 结构体
func Loads() {
	var jsonStr = `{"name": "dev", "age": 18}`
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
}

// JSON操作示例
func JsonDemo() {
	Dump()
	Dumps()
	Load()
	Loads()
}

// https://www.json.cn/ 验证json格式是否正确
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

// 将结构体转为json 进行序列化操作
func testStructToJson() {
	person := Person{
		Name:    "张三",
		Age:     18,
		Address: "北京",
	}
	//序列化操作 struct->json 需要传进去指针
	data, err := json.Marshal(&person)
	if err != nil {
		fmt.Println("json.Marshal failed,err:", err)
		return
	}
	fmt.Printf("person json:%s\n", data)
}

// map转为json格式
func testMapToJson() {
	map1 := make(map[string]interface{})
	map1["name"] = "李四"
	map1["age"] = 28
	map1["address"] = "上海"
	//序列化操作 map->json
	data, err := json.Marshal(map1)
	if err != nil {
		fmt.Println("json.Marshal failed,err:", err)
		return
	}
	fmt.Printf("map1 json:%s\n", data)
}
func main() {
	testStructToJson()
	testMapToJson()
}

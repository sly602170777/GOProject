package main

import (
	"fmt"
)

// 生成一个学生构造体，包含name,age字段
type Student struct {
	name string
	age  int
}

// 编写一个函数，可以判断输入多个参数是什么类型
func test(items ...interface{}) {
	for index, v := range items {
		index++
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, v)
		case float32, float64:
			fmt.Printf("第%v个参数是 float 类型，值是%v\n", index, v)
		case int, int8, int16, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, v)
		case string:
			fmt.Printf("第%v个参数是 字符串 类型，值是%v\n", index, v)
		//如果是其他类型则直接跳出判断
		default:
			fmt.Printf("第%v个参数是 其他类型，值是%v\n", index, v)
		}

	}
}

func main() {
	var s = Student{
		name: "张三",
		age:  18,
	}
	test(s, 10, 3.14, true, "hello")
}

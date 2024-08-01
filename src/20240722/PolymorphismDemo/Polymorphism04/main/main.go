package main

import (
	"fmt"
)

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
		default:
			fmt.Printf("第%v个参数是 未知类型%v\n", index, v)
		}

	}
}

type Student struct {
}

func main() {
	fmt.Println()
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300

	stu1 := Student{}
	stu2 := &Student{}
	test(n1, n2, n3, n4, name, address, stu1, stu2)

	//冒泡排序
	fmt.Println("冒泡排序")

	var arr = []int{1, 5, 8, 4, 2, 9, 3, 7, 6}
	fmt.Println("排序前：", arr)
	for i := 0; i < len(arr)-1; i++ { //控制比较轮数
		for j := 0; j < len(arr)-1-i; j++ { //控制每轮比较次数
			if arr[j] > arr[j+1] { //如果前一个数大于后一个数，则交换位置
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("排序后：", arr)

	//生成一个多表查询的SQL
	fmt.Println("生成一个多表查询的SQL")
	//日本語で出力　Hello world
	fmt.Println("日本語で出力　Hello world　")

}

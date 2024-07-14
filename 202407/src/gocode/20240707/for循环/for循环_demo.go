package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello world", i)
	}

	fmt.Println("①*********************")
	//字符串遍历传统方法
	var str string = "HELLO/WORLD"

	str2 := []rune(str) //将string转换成切片[]rune
	fmt.Println("①//将string转换成切片[]rune")
	for i := 0; i < len(str2); i++ {
		fmt.Printf("下标是：%d,值是：%c\n", i, str2[i])
	}

	for i := 0; i < len(str); i++ {
		fmt.Printf("下标是：%d,值是：%c\n", i, str[i])
	}
	//字符串遍历方法 for-range
	fmt.Println("③*********************")
	str = "APPLE WATCH"
	for _in, v := range str {
		fmt.Printf("下标是：%d,值是：%c\n", _in, v)
	}
}

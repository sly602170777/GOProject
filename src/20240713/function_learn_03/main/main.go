package main

import (
	"fmt"
	"strings"
)

// 闭包的使用 函数makeSuffix(suffix string)中的suffix变量
//组成的一个闭包，返回的值引用到这个变量

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	fun := makeSuffix(".jpg")

	fmt.Println("文件处理后是：", fun("winter"))  //winter.jpg
	fmt.Println("文件处理后是：", fun("abc.jpg")) //abc.jpg
}

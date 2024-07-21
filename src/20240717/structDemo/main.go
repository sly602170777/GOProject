package main

import (
	"fmt"
	master "struct"
)

func main() {
	var fruit master.Fruit
	fruit.Arr = []string{}        //申明切片方法一
	fruit.Arr = make([]string, 0) //申明切片方法二
	fmt.Println(fruit.Arr)

	fruit.Map1 = make(map[string]string, 2) //申明map
	fmt.Println(fruit.Map1)

	// 初始化一个整数指针变量，指向的值为0
	//var i3 *int = new(int)
	fruit.Ptr = new(int)    //申明整数指针变量
	fmt.Println(&fruit.Ptr) //0xc0000c8018
	fmt.Println(*fruit.Ptr) //0

}

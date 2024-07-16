package main

import (
	"fmt"
)

// 方式一：使用make进行申明切片
var slice []byte = make([]byte, 5, 10)

// 方式二：直接进行申明切片
var strSli []string = []string{"a", "a", "a"}

func main() {

	slice[1] = 20
	slice[2] = 30
	strSli[2] = "b"
	fmt.Println(slice)                   //[0 20 30 0 0]
	fmt.Println("slice的长度：", len(slice)) //slice的长度： 5
	fmt.Println("slice的容量：", cap(slice)) //slice的容量： 10
	fmt.Println("方式二：直接进行申明切片")
	fmt.Println(strSli)                   //[a a b]
	fmt.Println("slice的长度：", len(strSli)) //slice的长度： 3
	fmt.Println("slice的容量：", cap(strSli)) //slice的容量：3
}

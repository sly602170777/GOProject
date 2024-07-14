package main

import "fmt"

/**
*	init函数完成初始化工作、然后再执行main函数
	未来看到全局变量是先被初始化的
*
*
**/

var age = test()

func test() int {
	fmt.Println("this is test fun") //第一步执行
	return 99
}
func init() {
	fmt.Println("this is init fun") //第二步执行
}

func main() {
	fmt.Println("main fun。。。age", age) //第三步执行
}

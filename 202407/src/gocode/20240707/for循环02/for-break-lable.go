package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				break //默认会跳出最近的一个for循环
			}
			fmt.Println("j=", j)
		}

	}
	fmt.Println("--------")
	for i := 0; i < 4; i++ {
	lable1: //设置一个标签
		for j := 0; j < 5; j++ {
			if j == 2 {
				//break //默认会跳出最近的一个for循环
				break lable1 //将停止当前的标签
			}
			fmt.Println("j=", j)
		}

	}
}

package main

import (
	"fmt"
)

func test() {
	//使用defer recover 开捕获和处理异常
	defer func() {
		error := recover() //捕获异常
		if error != nil {
			fmt.Println("有个err:", error) //error err: runtime error: integer divide by zero
		} else {
			fmt.Println("没有err:", error) // error --》<nil>
		}
	}()
	num1 := 100
	num2 := 10
	res := num1 / num2
	fmt.Println("res :", res)
}

func main() {
	test()

	fmt.Println("main() end----") //使用defer 可以捕获异常后正常处理main以后的内容
}

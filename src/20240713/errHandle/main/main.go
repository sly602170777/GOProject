package main

import (
	user "20240713/errHandle/user"
	"errors"
	"fmt"
)

func test() {
	//使用defer recover 开捕获和处理异常
	defer func() {
		error := recover() //捕获异常
		//如果error不是空
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

// 自定义错误处理
func verifyUser(verifyUser user.User) (bool, error) {
	password, error := user.QueryUser(verifyUser.Name)
	if error != nil {
		return false, error
	}
	if password != verifyUser.Password {
		return false, errors.New("密码验证失败 ，请重新输入。。。")
	}
	return false, nil
}

// 实现一个user类
var user1 = user.User{Name: "jac", Password: "jack123"}

func main() {
	test()
	fmt.Println("main() end----") //使用defer 可以捕获异常后正常处理main以后的内容
	fmt.Println("main() user1----", user1)

	//进行验证user
	resule, err := verifyUser(user1)
	//如果有error 就输出错误并且终止执行
	if err != nil {
		panic(err)
	}
	fmt.Printf("resule----%v\n", resule)
}

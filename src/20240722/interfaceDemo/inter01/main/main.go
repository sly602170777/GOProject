package main

import (
	"fmt"
)

/*
*
接口申明

	1》InterSay
	2》InterHello
*/
type InterSay interface {
	say()
}
type InterHello interface {
	hello()
}

type Student struct {
	Name string
	Age  int
}

func (stu Student) say() {
	fmt.Println("This is student say()")
}

type integer int

func (i integer) say() {
	fmt.Println("This is integer say()", i)
}

type People struct {
}

func (p People) say() {
	fmt.Println("this is people say()")
}
func (p People) hello() {
	fmt.Println("this is people hello()")
}

func main() {
	fmt.Println("接口实现。。。")
	//构造体实现了接口
	var student Student
	var inter InterSay = student //结构体变量实现了say() 实现了Interface()
	student.say()
	inter.say() //接口指向一个实现该接口的自定义类型的变量

	//基本类型可以实现接口
	var i integer = 10
	var intTemp InterSay = i //基本类型的变量实现了say() 实现了Interface()
	intTemp.say()            //This is integer say() 10

	//一个变量实例可以实现多个接口
	var peo People
	var peoSay InterSay = peo
	peoSay.say() //构造体实现了接口 就实现 InterSay的方法say()
	var peoHello InterHello = peo
	peoHello.hello() //构造体实现了接口 就实现 InterHello的方法 hello()
}

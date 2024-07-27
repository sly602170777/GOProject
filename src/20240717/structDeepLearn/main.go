package main

import (
	"fmt"
)

type integer int

// 定义 struct的方法 接受者为值类型绑定的  传进来一个值拷贝
func (i integer) printNew() {
	i = 30
	fmt.Println("i=", i) //30 是通过值copy 到形参中
}

// 定义 struct方法 接受者为指针类型绑定的  传进来一个地址拷贝
func (i *integer) changePrint() {
	*i = 40
	fmt.Println("changePrint() i=", *i) //40 是通过地址copy 到形参中 将i=20-》i=40
}

// 定义 struct方法 接受者为指针类型绑定的  传进来一个地址拷贝 形参也是为指针类型进行地址拷贝
func (i *integer) changePrintHaveParam(val *int) {
	*i = 40 + 10
	*val = *val + 10
	fmt.Println("changePrintHaveParam() i=", *i)     // i= 50 是通过地址copy 到形参中 将i=20-》i=50
	fmt.Println("changePrintHaveParam() val=", *val) // val= 30
}

func main() {
	var i integer = 20
	i.printNew()
	fmt.Println("main() i=", i) //i= 20

	i.changePrint()
	fmt.Println("main() i=", i) //i= 40
	var k int = 20
	i.changePrintHaveParam(&k)
	fmt.Println("main() i=", i) //main() i= 50
	fmt.Println("main() k=", k) //main() k= 30

}

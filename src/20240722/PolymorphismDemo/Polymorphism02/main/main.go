package main

import (
	"fmt"
)

// 声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

// 让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
	name string
}

// 让Camera 实现   Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	fmt.Println(usbArr)
	usbArr[2].Start() //相机开始工作。
	usbArr[2].Stop()  //相机停止工作。。。

	var usbTemp interface{}
	var huawei = Phone{"华为"}
	usbTemp = huawei
	//类型断言 符合同一类型
	a := usbTemp.(Phone)
	fmt.Println(a) //{华为}
	//方法一：类型断言 不符合同一类型，不要报poinc 带检测的断言
	b, ok := usbTemp.(Camera)
	if ok {
		fmt.Println(ok)
	} else {
		fmt.Println("error")
	}
	//方法二：简洁版 类型断言 不符合同一类型，不要报poinc 带检测的断言
	if k, ok := usbTemp.(Camera); ok {
		fmt.Println("convert success", k)
	} else {
		fmt.Println("error", k)
	}
	fmt.Println("程序继续执行中。。。", b)
}

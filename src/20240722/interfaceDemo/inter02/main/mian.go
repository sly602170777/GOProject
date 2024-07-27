package main

import (
	"fmt"
)

type Usb interface {
	UsbSay()
}

type Student struct {
}

func (s *Student) UsbSay() {
	fmt.Println("Student ASay() 通过指针传递")
}

func main() {
	var stu Student
	//编译措误 会报stu没有实现Usb接口
	//应该是 var usb Usb = &stu
	//var usb Usb = stu
	var usb Usb = &stu
	usb.UsbSay()
	fmt.Println("-------")
}

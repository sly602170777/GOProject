package main

import (
	"fmt"
)

// 接口 通过拓展的方式 不破坏结构体继承关系，实际上是对继承的一种补充
// 1.声明一个构造体
type Monkey struct {
	Name string
}

// 继承Monkey结构体
type LittleMoney struct {
	Monkey
}

// 声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

// 将子类LittleMoney结构体与接口进行捆绑 实现接口的功能
func (m LittleMoney) BirdAble() {
	fmt.Println(m.Name, "：学会了飞行")
}

// 将子类LittleMoney结构体与接口进行捆绑 实现接口的功能
func (m LittleMoney) FishAble() {
	fmt.Println(m.Name, "：学会了下海")
}
func (m Monkey) climbable() {
	fmt.Println(m.Name, "：是猴子天生会爬树")
}

func main() {
	//方法一：创建一个LittleMonkey 实例
	var wukong LittleMoney
	wukong.Name = "悟空"
	//方法二：创建一个LittleMonkey 实例
	sixEarsMonkey := LittleMoney{
		Monkey{Name: "六耳猕猴"},
	}
	wukong.climbable()
	wukong.BirdAble()
	wukong.FishAble()

	sixEarsMonkey.climbable()
	sixEarsMonkey.BirdAble()
	sixEarsMonkey.FishAble()
}

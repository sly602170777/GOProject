package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.声明一个构造体
type Student struct {
	Name string
	Age  int
}

// 2.声明一个构造体切片
type StuentSlice []Student

// 3.实现Interface接口 Len()
func (ss StuentSlice) Len() int {
	return len(ss)
}

// func Sort(data Interface)  需要实现三个方法data.Len，data.Less和data.Swap
// Sort排序data。它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap。
// 本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）。
// 排序的方法 按年龄从小到大排序
func (ss StuentSlice) Less(i, j int) bool {
	return ss[i].Age < ss[j].Age
}
func (ss StuentSlice) Swap(i, j int) {
	//切片元素进行交换位置
	// temp := ss[i]
	// ss[i] = ss[j]
	// ss[j] = temp
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	//先定义一个切片
	var intSlice = []int{-1, 2, -5, 5, 9}
	//自动将切片元素从小到大排序
	sort.Ints(intSlice)
	fmt.Println("--对切片进行排序--")
	fmt.Println(intSlice)
	fmt.Println("--对结构体切片进行排序--")
	var students StuentSlice
	for i := 0; i < 10; i++ {
		student := Student{
			Name: fmt.Sprintf("学生|%d", rand.Intn(50)),
			Age:  rand.Intn(50),
		}
		students = append(students, student)
	}
	fmt.Println("排序前----")
	for _, v := range students {
		fmt.Println(v)
	}
	fmt.Println("排序后----")
	sort.Sort(students)
	//看看排序后的顺序
	for _, v := range students {
		fmt.Println(v)
	}
}

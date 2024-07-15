package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 函数内不能修改原有数组的值
func test01(arr [3]int) {
	arr[0] = 77
	fmt.Printf("test01() 函数内，arr值是%v\n", arr) //test01() 函数内，arr值是[77 5 9]
}

// 函数内修改原有数组的值
func test02(arr *[3]int) {
	(*arr)[0] = 77
	fmt.Printf("test02() 函数内，arr值是%v\n", arr) //test02() 函数内，arr值是&[77 5 9]
}

// 随机生成一个数组并将它进行反转打印
func test03() {
	var arr [5]int
	rand.Seed(time.Now().UnixMicro()) //使用纳秒时间戳，为了每次生成不一样的数组需要用seed
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) //随机生成100以内的数组
	}
	fmt.Println("随机生成的数组----反转之前", arr) // [84 15 94 40 85]

	var arrTemp [5]int
	for i := len(arr); i > 0; i-- {
		arrTemp[5-i] = arr[i-1]
	}
	fmt.Println("随机生成的数组----反转之后", arrTemp) // [85 40 94 15 84]
}

func main() {
	//创建一个int型的数组，没有长度限制
	array01 := [...]int{1, 4, 8}
	for i, v := range array01 {
		fmt.Printf("索引是%v--，值是%v\n", i, v)
	}

	array02 := [3]int{1, 5, 9}
	test01(array02)
	fmt.Printf("array02值是%v\n", array02) //array02值是[1 5 9]
	array03 := [...]int{1, 5, 9}
	test02(&array03)                     //传入的地址，通过指针传递 效率更高
	fmt.Printf("array02值是%v\n", array03) //array02值是[77 5 9]
	//调用随机生成一个数组的函数
	test03()
}

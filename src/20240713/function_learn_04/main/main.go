package main

import (
	"fmt"
)

// defer在函数中以先入后出的原则、在上边的后输出

func getSum(a int, b int) int {
	defer fmt.Printf("defer 01 : %v/n", a) // 再执行 01  a->2
	defer fmt.Printf("defer 02: %v/n", b)  // 然后执行 02 b->4
	a++
	b++
	fmt.Println("没有使用defer") // 首先没有使用defer
	return a + b

}

func main() {
	result := getSum(2, 4)
	fmt.Println("getSum 结果：", result) // 最后执行 8
}

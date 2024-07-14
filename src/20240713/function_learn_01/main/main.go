package main

import "fmt"

/*
	递归调用的重要原则：
		1》执行一个函数时，就创建一个新的受保护的独立空间（栈区）
		2》函数局部变量是独立的，不会相互影响
		3》递归必须向退出递归条件逼近 否则就成了死循环
		4》当一个函数执行完毕或者遇到return就会返回,
		 遵守谁调用就将结果返回给谁，同时当函数执行完毕或者返回时，
		 该函数本身也会被系统销毁
*/

func test(n int) {
	if n > 2 {
		n--
		fmt.Println("n-- ", n)
		test(n)
	}
	fmt.Println("n= ", n)
}

func test2(n int) {
	if n > 2 {
		n--
		fmt.Println("n-- ", n)
		test2(n)
	} else {
		fmt.Println("n= ", n)
	}
}

//n1 就是 *int 类型
func test3(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("n1 ---- ", n1)         //n1 地址 0xc000016160
	fmt.Println("&n1 ---- ", &n1)       //%n1 地址 0xc00000a030
	fmt.Println("n1 function () ", *n1) //*n1 是取值 30
}
func getSum(a int, b int) int {
	return a + b
}

//此处传的是一个函数 作为一个值变量
func getSunRestlt(getSum int) int {
	a := getSum + 1
	return a
}

//首先申明一个type 类型的函数
type myFunType func(int, int) int

//然后将type 类型的函数作为值进行传递
func myFunTypeShow(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//支持对函数返回值的命名 不需要在return中添加返回值 可以返回两个值
func getSumAndSub(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

//支持可变参数，有多个参数
//形参列表中有可变参数，必须要放在最后
func sum(n1 int, args ...int) int {
	sum := n1
	//forr循环下标为：_index 值为：v 或者 args[_index]
	for _index, v := range args {
		fmt.Printf("forr循环的下标：%v forr循环的值：%v\n", _index, args[_index])
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("START ")
	//递归用法一：
	test(4) //  2,2,3
	//递归用法二：
	test2(4) //  2
	fmt.Println("END ")
	//通过指针传递实参
	n3 := 20
	//传入的参数是一个地址  &n3 是一个n3的地址，通过函数test3()将n3的地址对应的值改变
	test3(&n3)
	fmt.Println("n3 main的值 ", n3) //30
	fmt.Println("函数也是一种数据类型---")
	a := getSum
	fmt.Printf("a的数据类型%T：\n", a)
	fmt.Printf("getSum的数据类型%T：\n", getSum)

	result := a(3, 5)
	fmt.Println("result的值：", result)
	//将函数作为一个形参
	result01 := getSunRestlt(a(3, 5))
	fmt.Println("getSunRestlt的值：", result01)

	//自定义数据类型
	type myInt int32
	var num001 myInt

	num002 := 2147
	//需要对num002进行显示转换 在这里 int 和 myInt是两个不同的数据类型
	num001 = myInt(num002)
	fmt.Printf("num001: %v \n", num001) //2147
	fmt.Printf("num002: %v \n", num002) //2147

	result003 := myFunTypeShow(getSum, 10, 20)
	fmt.Println("type类型的函数作为数据类型传递的值：", result003)

	//需要申明两个变量接受函数的返回值
	result004, result005 := getSumAndSub(1, 9)
	fmt.Println("result005的sum值：", result004)
	fmt.Println("result005的sub值：", result005)

	result006 := sum(1, 4, 6)
	fmt.Println("result006的sum值：", result006)
}

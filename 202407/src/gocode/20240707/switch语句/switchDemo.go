package main

import "fmt"

// 第一段代码 Hello World
/* func main() {
	var days int = 97
	var weeks int = days / 7
	fmt.Printf("%d天是第%d个星期", days, weeks)

	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Printf("%d岁已经成年", age)
	} else {
		fmt.Printf("%d岁还未成年", age)
	}
} */

//switch 入门实例一
/* func main() {
	var key byte
	fmt.Println("a,b,c,d...")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("a..周一")
	case 'b':
		fmt.Println("b..周2")
	default:
		fmt.Println("nothing")
	}
}
*/

func test(str byte) byte {
	return str + 1
}

//switch 入门实例二 ，
/* 注意：
控制变量可以是一个方法返回值
并且需要返回值的数据类型和控制变量保持一致
*/
func main() {
	/* 	var key byte
	   	fmt.Println("a,b,c,d...")
	   	fmt.Scanf("%c", &key)
	   	switch test(key) {
	   	case 'a':
	   		fmt.Println("a..周一")
	   	case 'b':
	   		fmt.Println("b..周2")
	   	default:
	   		fmt.Println("nothing")
	   	} */

	//switch 入门实例三 ，
	/* 注意：
	1>控制变量可以是多个表达式
	2>并且需要返回值的数据类型和控制变量保持一致
	3>case 后不需要添加break
	4>default 不是必须的 可以没有
	*/
	//switchMothed_demo2()

	//switch 后不接表达式 等同于if - else 分支的用法
	//switchMothed_demo3()

	//switch 可以用穿透到下一个case
	//fallthrough 关键字
	switchMothed_demo4()
}

/* func switchMothed_demo2() {
	fmt.Println("可以跟多个表达式-----")
	var key int = 7
	var key2 int = 20

	switch key {
	case key2, 10, 9: //key 与case 后任意一个表达式进行比较
		fmt.Println("OK")
	default:
		fmt.Println("nothing")
	}
} */

/* func switchMothed_demo3() {
	fmt.Println("可以配合if else 使用-----")

	var score int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&score)

	switch {
	case score >= 90:
		fmt.Println("成绩优秀")
	case score > 60 && score < 90:
		fmt.Println("成绩良好")
	default:
		fmt.Println("成绩不及格")
	}
} */

func switchMothed_demo4() {
	num := 1
	switch num {
	case 1:
		fmt.Print("I ")
		fallthrough
	case 2:
		fmt.Print("am ")
		fallthrough
	case 3:
		fmt.Println("yyh-gl.")
		// fallthrough // 次の節がなければコンパイルエラー
	}
}

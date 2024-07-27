package main

import (
	"encoding/json"
	"fmt"
	master "struct"
)

func getPeopleInf() string {
	var people master.People //{"No001", "tom", 20}
	people.ID = "No002"
	people.Age = 20
	people.Name = "tom"

	//bob := master.People{"No001", "tom", 20} 不推荐使用
	//bob := master.People{ID: "No001", Name: "tom", Age: 20} -> 推荐使用
	jsonStr, erro := json.Marshal(people) //转化为json格式的字符串
	if erro != nil {
		fmt.Println("出现了错误")
	}
	fmt.Println("jsonStr:", string(jsonStr))
	return string(jsonStr) //然后通过将json转换为string进行输出
}

func main() {
	var fruit master.Fruit
	fruit.Arr = []string{}        //申明切片方法一
	fruit.Arr = make([]string, 1) //申明切片方法二 首先要make()
	fruit.Arr[0] = "No001"
	arr1 := "No002"
	fruit.Arr = append(fruit.Arr, arr1)
	fmt.Println(fruit.Arr)

	fruit.Map1 = make(map[string]string, 2) //申明map 首先要make()
	fruit.Map1["No"] = "No110"
	fruit.Map1["Name"] = "hero"
	fmt.Println(fruit.Map1)

	// 初始化一个整数指针变量，指向的值为0
	//var i3 *int = new(int)
	fruit.Ptr = new(int)    //申明整数指针变量
	fmt.Println(&fruit.Ptr) //0xc0000c8018
	*fruit.Ptr = 123
	fmt.Println(*fruit.Ptr) //123

	//strut 是值传递，不同的对象值是不相互影响
	var order1 master.Order
	var order2 master.Order
	order1.OrderID = 123
	order1.CustomerName = "smith"
	order2 = order1 //只是值进行了copy order2开辟了一个新的地址空间 地址没有相互传递
	order2.CustomerName = "Jim"
	if &order1 == &order2 { //&order1 == &order2  -》false
		fmt.Println("OK")
	}
	fmt.Println("order1:", order1) //order1: {123  0 smith}
	fmt.Println("order2:", order2) //order2: {123  0 Jim}

	var order3 master.Order
	order3.OrderID = 234
	order3.CustomerName = "jan"
	var order4 *master.Order = &order3  //将order3的地址交给order4,
	fmt.Println((*order4).OrderID)      //234
	fmt.Println((*order4).CustomerName) //jan
	order4.OrderID = 999
	order4.CustomerName = "king"

	//分析order3与order4 地址与值的分析
	fmt.Printf("order3的地址%p,order3的值%p\n", &order3, &order3)
	fmt.Printf("order4的地址%p,order4的值%p\n", &order4, &order4)
	fmt.Println("order3的与order4的值是否相等:", order3.CustomerName == order4.CustomerName)
	//order3.CustomerName=king,order4.CustomerName=king
	fmt.Printf("order3.CustomerName=%v,(*order4).CustomerName=%v\n", order3.CustomerName, (*order4).CustomerName)
	//order3.CustomerName=king,order4.CustomerName=king
	//(*order4).CustomerName √
	//*order4.CustomerName   ×   因为.的优先级比*要高，不能对一个具体的值再进行取值
	fmt.Printf("order3.CustomerName=%v,(*order4).CustomerName=%v\n", order3.CustomerName, (*order4).CustomerName)

	str := getPeopleInf()
	//未使用struct的tag标签时 {"ID":"No002","Name":"tom","Age":20}
	fmt.Printf("转义后的json格式%v\n", str) //{"id":"No002","name":"tom","age":20}

	//为model包添加一个NewStudent方法 返回一个指针型的student
	student := master.NewStudent("termp", 78)
	fmt.Println("工厂模式下的student：", *student) // {termp 78}
	//fmt.Println(student.name) 当name首字母小写时不能直接访问需要进行函数处理
	fmt.Println(student.GetStudentName()) //termp
	fmt.Println(student.Age)

	//使用匿名构造体 实现继承 小学生
	prisut := &master.PrimaryStudent{}
	//指定匿名构造体名字来给属性赋值
	prisut.Student.Name = "Tom"
	prisut.Student.Age = 12
	prisut.Testing()
	prisut.Student.SetScore(90)
	prisut.ShowInfo()

	//使用匿名构造体 实现继承 高中生
	higsut := &master.HighStudent{}
	//指定匿名构造体名字来给属性赋值
	higsut.Student.Name = "Mary"
	higsut.Student.Age = 42
	higsut.Testing()
	higsut.Student.SetScore(80)
	higsut.ShowInfo()

}

package main

import (
	"20240730/model"
	"20240730/service"
	"fmt"
)

type customerView struct {
	key             string                   // 接收用户输入的指令
	loop            bool                     // 表示是否循环显示菜单
	customerService *service.CustomerService // 声明一个字段，用于调用CustomerService的方法
}

// 显示所有客户的信息
func (cv *customerView) list() {
	//首先，获取到当前所有的客户信息(在切片中)
	customers := cv.customerService.List()
	// 遍历，并显示
	fmt.Println("----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t年龄\t地址\t电话\t邮箱")
	for _, customer := range customers {
		fmt.Println(customer.GetInfo())
		fmt.Println()
	}
}

// 得到用户的输入，信息构建新的用户，并完成添加客户
func (cv *customerView) add() {
	fmt.Println("----------------添加客户-----------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("地址：")
	address := ""
	fmt.Scanln(&address)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer
	customer := model.NewCustomer2(name, age, address, phone, email)
	//调用service的add方法
	flag := cv.customerService.Add(*customer)
	if flag {
		fmt.Println("----------------添加完成-----------------")
		fmt.Println()
	} else {
		fmt.Println("----------------添加失败-----------------")
		fmt.Println()
	}
}

// 退出系统
func (cs *customerView) exit() {
	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scan(&cs.key)
		if cs.key == "Y" || cs.key == "y" || cs.key == "N" || cs.key == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入(Y/N)：")
	}
	if cs.key == "Y" || cs.key == "y" {
		cs.loop = false
		fmt.Println("退出系统")

	}
}

// 删除客户
func (cv *customerView) delete() {
	fmt.Println("----------------删除客户-----------------")
	fmt.Println("请输入客户id，-1退出：")
	id := -1
	fmt.Scan(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N)：")
	fmt.Scan(&cv.key) //&cv.keyは構造体のメソッドの引数

	if cv.key == "Y" || cv.key == "y" {
		flag := cv.customerService.Delete(id)
		if flag {
			fmt.Println("----------------删除完成-----------------")
			fmt.Println()
		} else {
			fmt.Println("----------------删除失败-----------------")
			fmt.Println()
		}
	}
}

// 更新客户
func (cv *customerView) update() {
	fmt.Println("----------------更新客户-----------------")
	fmt.Println("请输入客户id，-1退出：")
	id := -1
	fmt.Scanln(&id)
	id, index := cv.customerService.FindById(id)

	if id == -1 && index == -1 {
		fmt.Printf("ID:%v客户不存在请重新输入", id)
		fmt.Println()
		return
	}
	fmt.Println("确认是否删除(Y/N)：")
	fmt.Scanln(&cv.key) //&cv.keyは構造体のメソッドの引数

	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("地址：")
	address := ""
	fmt.Scanln(&address)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer
	customer := model.NewCustomer2(name, age, address, phone, email)

	if cv.key == "Y" || cv.key == "y" {
		flag := cv.customerService.Update(id, *customer)
		if flag {
			fmt.Println("----------------更新完成-----------------")
			fmt.Println()
		} else {
			fmt.Println("----------------更新失败-----------------")
			fmt.Println()
		}
	}
}

// 构造体的方法 显示主菜单 構造体と関数を紐つける
// ①：構造体のメソッドは、構造体のインスタンスを引数として受け取ります。
// ②：構造体のメソッドは、構造体のフィールドにアクセスして、その値を操作できます。
/* func (cv *customerView) mainMenu() {

} */
func (cv *customerView) showMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
	}
}

func main() {
	customerView := customerView{key: "", loop: true}
	customerView.customerService = service.NewCustomerService()
	//构造体的显示主菜单方法
	customerView.showMenu()
}

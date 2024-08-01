package service

import (
	"20240730/model"
	"fmt"
)

// 该CustomerService， 完成对Customer的操作,包括
// 增删改查
type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

// 初始化CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{
		customers:   make([]model.Customer, 0, 10),
		customerNum: 1,
	}
	//var customer *model.Customer
	customer := model.NewCustomer(1, "张三", 20, "北京", "123456789", "zs@sohu.com")
	customerService.customers = append(customerService.customers, *customer)

	return customerService
}

// 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers

}

// 添加客户到customers切片
func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true

}

/**
* 根据id删除客户
* @params {int} id 客户的id
* @returns:{bool} 删除成功返回true，否则返回false
 */
func (cs *CustomerService) Delete(id int) bool {
	id, index := cs.FindById(id)

	fmt.Printf("id=%v,index=%v", id, index)
	if id == -1 {
		return false
	} else {
		fmt.Println("前---cs.customers=", cs.customers)
		//append(a[:index], a[index+1:]...) 删除切片中第index个的元素
		cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
		fmt.Println("后---cs.customers=", cs.customers)
		return true
	}
}

func (cs *CustomerService) Update(id int, customer model.Customer) bool {
	fmt.Println("cs.customerNum=", cs.customerNum)
	fmt.Println("id=", id)
	if cs.customerNum == id {
		id, index := cs.FindById(id)
		fmt.Println("id=", id)
		cs.customers[index].Name = customer.Name
		cs.customers[index].Age = customer.Age
		cs.customers[index].Address = customer.Address
		cs.customers[index].Phone = customer.Phone
		cs.customers[index].Email = customer.Email
		return true
	}
	return false
}

/**
* 根据id查找客户在切片中的位置
* @params {int} id 客户的id
* @returns:{int, int} 返回客户在切片中的id和索引，如果不存在，则返回-1, -1
 */
func (cs *CustomerService) FindById(id int) (int, int) {
	for index, v := range cs.customers {
		if v.Id == id {
			return id, index
		}
	}
	return -1, -1
}

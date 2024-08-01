package model

import (
	"fmt"
)

// 声明一个Customer结构体，表示一个客户信息
type Customer struct {
	Id      int    //编号
	Name    string //姓名
	Age     int    //年龄
	Address string //地址
	Phone   string //手机号
	Email   string //邮箱
}

// 使用工厂模式，返回一个Customer实例
func NewCustomer(id int, name string, age int, address string, phone string, email string) *Customer {
	return &Customer{
		Id:      id,
		Name:    name,
		Age:     age,
		Address: address,
		Phone:   phone,
		Email:   email,
	}
}

// 添加一个不带ID的工厂方法
func NewCustomer2(name string, age int, address string, phone string, email string) *Customer {
	return &Customer{
		Name:    name,
		Age:     age,
		Address: address,
		Phone:   phone,
		Email:   email,
	}

}

// 返回用户的信息,格式化的字符串
func (customer Customer) GetInfo() string {
	info := fmt.Sprintf("id=%d\tname=%s\tage=%d\taddress=%s\tphone=%s\temail=%s\t", customer.Id, customer.Name, customer.Age, customer.Address, customer.Phone, customer.Email)

	return info

}

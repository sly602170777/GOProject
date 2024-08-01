package main

import (
	"20240722/PolymorphismDemo/Polymorphism06/model"
	"fmt"
)

func main() {
	account := model.NewAccount("1234567890", "123", 10000.0)
	fmt.Println(account)
}

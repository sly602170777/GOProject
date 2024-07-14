package main

import (
	db "20240713/package_learn/DB01"
	unit "20240713/package_learn/unitTool"
	"fmt"
)

func main() {
	db.GetDB01()
	unit.GetUnit()
	fmt.Println("this is main")
}

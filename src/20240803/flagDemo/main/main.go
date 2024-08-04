package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// 定义命令行参数
	var name string
	var age int
	var married bool
	var delay time.Duration // 声明一个命令行参数 格式：00:24:32.460

	timeDur, _ := time.ParseDuration("1472.46s")
	fmt.Println(timeDur) // Output: 24m32.46s

	// 解析命令行参数
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", timeDur, "延迟的时间间隔")

	// 解析命令行参数
	flag.Parse()

	//输出结果
	fmt.Printf("name=%v age=%v married=%v delay=%v",
		name, age, married, delay)

	//首先build，然后执行go run main.go -name=李四 -age=20 -married=false -d=1h30m
	//执行build命令并重命名为myApp  go build -o myApp.exe main.go
	//myApp.exe -name 东京 -age 10 -married true -d 00:24:32.460
}

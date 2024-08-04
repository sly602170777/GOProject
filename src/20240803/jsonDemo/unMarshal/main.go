package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

// 将Person结构体进行序列化
func structUnmarshal() string {
	person := Person{
		Name: "Alice", Age: 30, City: "New York",
	}
	personConvert, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	return string(personConvert)
}

// 将map进行序列化
func testMapToJson() string {
	m := make(map[string]interface{})
	m["name"] = "mary"
	m["age"] = 18
	m["city"] = "东北"
	mConvert, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	return string(mConvert)
}

// 将slice进行序列化
func testSliceToJson() string {
	s := make([]map[string]interface{}, 0)
	//方法一 ：对切片的第一个元素进行初始化
	// s[0]["name"] = "张三"
	// s[0]["age"] = 20
	// s[0]["city"] = "河南"

	//方法二 ：对切片的第一个元素使用make进行初始化
	s1 := make(map[string]interface{}, 3)
	s1["name"] = "张三"
	s1["age"] = 20
	s1["city"] = "河南"

	s = append(s, s1)

	// s[1] = make(map[string]interface{}, 3)
	// s[1]["name"] = "李四"
	// s[1]["age"] = 22
	// s[1]["city"] = "四川"

	s2 := make(map[string]interface{}, 3)
	s2["name"] = "李四"
	s2["age"] = 30
	s2["city"] = "信阳"

	//m := map[string]int{}
	s = append(s, s2)

	sConvert, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//fmt.Println("sConvert--", sConvert)
	return string(sConvert)

}

func main() {
	//jsonStr := `{"name":"Alice","age":30,"city":"New York"}`
	//使用函数返回值为string类型时，调用进行反序列化不需要转义
	jsonStr := structUnmarshal()
	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//反序列化后的结果封装在person
	fmt.Printf("反序列化后的person:name=%v,age=%d,city=%s\n", person.Name, person.Age, person.City)

	mapStr := testMapToJson()
	//定义一个map，类型为string，interface{}
	var m map[string]interface{}
	//反序列化
	//注意：反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err = json.Unmarshal([]byte(mapStr), &m)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//反序列化后的结果封装在m
	fmt.Printf("反序列化后的map:name=%v,age=%v,city=%s\n", m["name"], m["age"], m["city"])

	sliceStr := testSliceToJson()
	//fmt.Println("sliceStr--", sliceStr)
	//定义一个map，类型为string，interface{}
	var s []map[string]interface{}
	//反序列化
	//注意：反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err = json.Unmarshal([]byte(sliceStr), &s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//反序列化后的结果封装在s
	for _, v := range s {
		fmt.Printf("反序列化后的slice:name=%v,age=%v,city=%s\n", v["name"], v["age"], v["city"])

	}
}

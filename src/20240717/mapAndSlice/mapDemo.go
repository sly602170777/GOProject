package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 先申明一个切片map
var masters = make([]map[string]string, 3)

//masters = make([]map[string]string, 2)

// 向切片map里追加数据
func showMastersINFO() []map[string]string {
	//再对每个master进行申明
	masters[0] = make(map[string]string)
	//申明后的master进行key value 赋值
	masters[0]["ID"] = "No02"
	masters[0]["name"] = "mary"
	masters[0]["age"] = "10"

	masters[1] = make(map[string]string)
	masters[1]["ID"] = "No03"
	masters[1]["name"] = "tom"
	masters[1]["age"] = "120"

	masters[2] = make(map[string]string)
	masters[2]["ID"] = "No01"
	masters[2]["name"] = "jams"
	masters[2]["age"] = "19"

	return masters
}

// 更新map可以看到 map是指针 引用传递
func updateMaster() []map[string]string {
	//再对每个master进行申明
	masters[0] = make(map[string]string)
	//申明后的master进行key value 赋值
	masters[0]["ID"] = "No02"
	masters[0]["name"] = "mary"
	masters[0]["age"] = "80"
	return masters
}

// 当超过切片的容量时 需要用到append进行添加对象
func addMaster() map[string]string {
	//再对每个master进行申明
	masterAdd := make(map[string]string, 3)
	masterAdd["ID"] = "No04"
	masterAdd["name"] = "kim"
	masterAdd["age"] = "5"
	return masterAdd
}

// 通过年龄对master进行排序
func sortMasterByAge() []map[string]string {
	//slicesAge := make([]int, 10) 申明切片方法一
	slicesAge := []int{} //申明切片方法二

	for k, value := range masters {
		intAge, _ := strconv.Atoi(value["age"])
		slicesAge = append(slicesAge, intAge)
		fmt.Println("k--", k)
	}
	sort.Ints(slicesAge)
	fmt.Println("排完序后：", slicesAge)

	masterTemp := make([]map[string]string, 3)
	for i := 0; i < len(slicesAge); i++ {
		//将int 转为string
		str := fmt.Sprintf("%d", slicesAge[i])
		//fmt.Println("first level", intAge)
		for j := 0; j < len(masters); j++ {
			//mt.Println("second level value", len(masters))
			//将string 转为 int
			//intAge, _ := strconv.Atoi(masters[j]["age"])
			//if intAge == slicesAge[i] {
			if str == masters[j]["age"] {
				//fmt.Println("third level")
				fmt.Printf("j=%v,masterのID:%v", j, masters[j]["ID"])
				masterTemp = append(masterTemp, masters[j])
			}
		}
	}
	return masterTemp
}

func main() {
	var masterALL = make([]map[string]string, 3)
	//展示所有的master
	masterALL = showMastersINFO()
	for _, v := range masterALL {
		fmt.Printf("add master length : %v", len(masterALL))
		fmt.Println("函数展示所有的master---", v) //[map[age:80 name:mary] map[age:20 name:tom] map[age:5 name:kim]]
	}
	fmt.Println("使用函数接受一个切片map:", masterALL) // [map[age:10 name:mary] map[age:20 name:tom]]

	//修改master
	masterALL = updateMaster()
	for _, v := range masterALL {
		fmt.Printf("update master length : %v", len(masterALL))
		fmt.Println("更新后一个切片map:", v) //[map[age:80 name:mary] map[age:20 name:tom]]
	}

	//通过函数添加一个，append进行添加对象
	masters = append(masters, addMaster())
	for _, v := range masters {
		fmt.Printf("add master length : %v", len(masters))
		fmt.Println("函数添加一个masters---", v) //[map[age:80 name:mary] map[age:20 name:tom] map[age:5 name:kim]]
	}

	//排序后的master,需要注意的是整形转string 需要用到方法  str := fmt.Sprintf("%d", slicesAge[i])
	masters = sortMasterByAge()
	for _, v := range masters {
		if v != nil {
			fmt.Printf("sort master length : %v\n", len(masters))
			fmt.Println("排序后的masters---", v)
		}
	}

}

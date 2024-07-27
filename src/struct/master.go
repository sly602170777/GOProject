package master

import "fmt"

type User struct {
	Name     string
	Password string
	Age      int
	Address  string
}

type Order struct {
	OrderID      int
	ProductName  string
	Quantity     int
	CustomerName string
}

type Fruit struct {
	Name  string
	Price int
	Ptr   *int              //<- ポインタ定義
	Arr   []string          //切片
	Map1  map[string]string //map
}

/** JSONデコード用に構造体定義 **/
type People struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student struct {
	Name  string
	Age   int
	Score int
}

// 小学学生
type PrimaryStudent struct {
	Student //嵌入一个匿名的结构体
}

//小学生所特有方法
func (p *PrimaryStudent) Testing() {
	fmt.Println("小学生在考试。。。")
}

type HighStudent struct {
	Student //嵌入一个匿名的结构体
}

//高中生特有方法
func (p *HighStudent) Testing() {
	fmt.Println("高中学生在考试。。。")
}

//结构体共通方法，即子类所共有的方法
func (student *Student) ShowInfo() {
	fmt.Printf("学生名字%v，学生年龄%v,学生成绩%v\n", student.Name, student.Age, student.Score)
}

//结构体共通方法，即子类所共有的方法
func (student *Student) SetScore(score int) {
	student.Score = score
}

/** 当struts首字母是小写时，可以使用工厂模式进行申明 **/
//工厂模式来解决此问题
type student struct {
	name string
	Age  int
}

// 工厂模式的申明结构体
func NewStudent(name string, age int) *student {
	return &student{
		name: name,
		Age:  age,
	}
}

// 结构体的元素小写时
// 可以通过定义 struct方法 接受者为指针类型绑定的  传进来一个地址拷贝
func (s *student) GetStudentName() string {
	return s.name
}

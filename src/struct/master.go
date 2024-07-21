package master

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
	Ptr   *int              //指针
	Arr   []string          //切片
	Map1  map[string]string //map
}

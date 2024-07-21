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

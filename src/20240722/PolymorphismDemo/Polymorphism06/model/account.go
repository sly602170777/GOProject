package model

//定义一个结构体account
type account struct {
	accountNo string
	pwd       string
	balance   float64
}

//工厂模式的函数-构造函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

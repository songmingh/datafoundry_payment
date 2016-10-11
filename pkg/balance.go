package pkg

type BalanceAgent service

// type Balance struct {
// 	apiBalance
// }

type Balance apiBalance

func (agent *BalanceAgent) Get() *Balance {
	balance := &Balance{
		Balance: 50000.89,
		Status:  "active",
	}
	return balance
}

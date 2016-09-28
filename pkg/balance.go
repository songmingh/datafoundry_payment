package pkg

type BalanceAgent service
type Balance struct {
	Balance float32 `json:"balance"`
	Status  string  `json:"status,omitempty"`
}

func (u *BalanceAgent) Get() *Balance {
	balance := &Balance{
		Balance: 50000.89,
		Status:  "active",
	}
	return balance
}

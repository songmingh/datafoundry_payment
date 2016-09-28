package pkg

type RechargeAgent service

type Recharge struct {
	Amount  float32 `json:"amount"`
	Project string  `json:"namespace,omitempty"`
}

func (*RechargeAgent) Get() *Balance {
	balance := &Balance{
		Balance: 6000.89,
		Status:  "active",
	}
	return balance
}

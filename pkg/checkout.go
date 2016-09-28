package pkg

type CheckoutAgent service

type Checkout struct {
	PlanId  string `json:"plan_id"`
	Project string `json:"namespace,omitempty"`
	Region  string `json:"region"` //need it?
}

func (*CheckoutAgent) Get() *Balance {

	balance := &Balance{
		Balance: 3000.01,
		Status:  "active",
	}
	return balance
}

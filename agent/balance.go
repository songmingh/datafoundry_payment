package agent

import (
	"github.com/asiainfoLDP/datafoundry_payment/api"
)

type BalanceAgent service

func (u *BalanceAgent) Get() *api.Balance {
	balance := &api.Balance{}
	return balance
}

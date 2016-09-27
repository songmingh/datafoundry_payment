package agent

import (
	"github.com/asiainfoLDP/datafoundry_payment/api"
)

type AmountAgent service

func (u *AmountAgent) Get() *api.Amount {
	amount := &api.Amount{}
	return amount
}

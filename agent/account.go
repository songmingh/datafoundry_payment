package agent

import (
	"github.com/asiainfoLDP/datafoundry_payment/api"
)

type AccountAgent service

func (u *AccountAgent) Get() *api.Account {
	account := &api.Account{}
	return account
}

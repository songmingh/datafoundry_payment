package pkg

import (
	"net/http"

	"github.com/asiainfoLDP/datafoundry_payment/api"
	"github.com/asiainfoLDP/datafoundry_payment/fake"
)

type AccountAgent service

func (u *AccountAgent) Get(r *http.Request) *api.Account {
	account := fake.Account(r)
	return account
}

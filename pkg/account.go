package pkg

import (
	"net/http"
)

type AccountAgent service

type Account struct {
	Purchased bool    `json:"purchased"`
	Notify    bool    `json:"notification"`
	Plans     []Plan  `json:"subscriptions,omitempty"`
	Status    string  `json:"status"`
	Balance   Balance `json:"balance"`
}

func (u *AccountAgent) Get(r *http.Request) *Account {
	account := fakeAccount(r)
	return account
}

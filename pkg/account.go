package pkg

import (
	"net/http"

	"github.com/zonesan/clog"
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
	r.ParseForm()

	project := r.FormValue("project")

	clog.Debug(project)

	account := fakeAccount(r)
	return account
}

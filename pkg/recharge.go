package pkg

import (
	"net/http"
	"net/url"

	"github.com/zonesan/clog"
)

type RechargeAgent struct {
	*Agent
	BaseURL *url.URL
}

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

func (agent *RechargeAgent) Create(r *http.Request, recharge *Recharge) (*Balance, error) {

	urlStr := "/charge/v1/recharge"

	balance := new(Balance)

	if err := doRequest(agent, r, "POST", urlStr, recharge, balance); err != nil {
		clog.Error(err)

		return nil, err
	}

	return balance, nil

}

func (agent *RechargeAgent) Url() *url.URL {
	return agent.BaseURL
}

func (agent *RechargeAgent) Instance() *Agent {
	return agent.Agent
}

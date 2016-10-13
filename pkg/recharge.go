package pkg

import (
	"encoding/json"
	"io"
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

func (agent *RechargeAgent) Notification(r *http.Request) error {
	urlStr := "/charge/v1/aipaycallback"
	var reqbody interface{}

	if r.URL.RawQuery != "" {
		urlStr += "?" + r.URL.RawQuery
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	u := agent.BaseURL.ResolveReference(rel)

	err = json.NewDecoder(r.Body).Decode(reqbody)
	if err == io.EOF {
		err = nil // ignore EOF errors caused by empty response body
	} else {
		return err
	}

	req, err := agent.NewRequest("POST", u.String(), reqbody)
	if err != nil {
		return err
	}

	response := new(RemoteResponse)

	if err := agent.Do(req, response); err != nil {
		return err
	}

	return nil
}

func (agent *RechargeAgent) Url() *url.URL {
	return agent.BaseURL
}

func (agent *RechargeAgent) Instance() *Agent {
	return agent.Agent
}

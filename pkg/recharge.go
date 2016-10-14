package pkg

import (
	"encoding/json"
	"io"
	// "io/ioutil"
	"net/http"
	"net/url"

	"github.com/zonesan/clog"
)

type RechargeAgent struct {
	*Agent
	BaseURL *url.URL
}

type Recharge struct {
	Amount  float64 `json:"amount,omitempty"`
	Project string  `json:"namespace,omitempty"`
}

type HongPay apiRechargePayload

// func (*RechargeAgent) Get() *Balance {
// 	balance := &Balance{
// 		Balance: 6000.89,
// 		Status:  "active",
// 	}
// 	return balance
// }

func (agent *RechargeAgent) Create(r *http.Request, recharge *Recharge) (*HongPay, error) {

	urlStr := "/charge/v1/recharge"

	hongpay := new(HongPay)

	if err := doRequest(agent, r, "POST", urlStr, recharge, hongpay); err != nil {
		clog.Error(err)

		return nil, err
	}
	clog.Debug(hongpay.Payloads)
	return hongpay, nil

}

func (agent *RechargeAgent) Notification(r *http.Request) error {
	urlStr := "/charge/v1/aipaycallback"
	reqbody := new(json.RawMessage)

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
		clog.Warn("empty request body....")
		err = nil // ignore EOF errors caused by empty response body
	} else {
		if err != nil {
			clog.Error(err)
			return err
		}
	}

	// b, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	return err
	// }
	// clog.Debug("Request Body:", string(b))
	// if err := json.Unmarshal(b, &reqbody); err != nil {
	// 	clog.Error(err)
	// 	return err
	// }

	req, err := agent.NewRequest("POST", u.String(), reqbody)
	if err != nil {
		clog.Error(err)
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

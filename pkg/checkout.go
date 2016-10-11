package pkg

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/zonesan/clog"
)

type CheckoutAgent struct {
	*Agent
	BaseURL *url.URL
}

type Checkout struct {
	PlanId  string `json:"plan_id"`
	Project string `json:"namespace,omitempty"`
	Region  string `json:"region"` //need it?
	OrderID string `json:"orderid,omitempty"`
}

func (agent *CheckoutAgent) Get() *Balance {

	balance := &Balance{
		Balance: 3000.01,
		Status:  "active",
	}
	return balance
}

func (agent *CheckoutAgent) ListOrders(r *http.Request) (*[]apiPurchaseOrder, error) {
	urlStr := "/usageapi/v1/orders"

	if r.URL.RawQuery != "" {
		urlStr += "?" + r.URL.RawQuery
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := agent.BaseURL.ResolveReference(rel)

	req, err := agent.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	response := new(RemoteListResponse)
	if err := agent.Do(req, response); err != nil {
		clog.Error(err)
		return nil, err
	}

	orders := []apiPurchaseOrder{}

	if err := json.Unmarshal([]byte(response.Data), &orders); err != nil {
		clog.Error(err)
		return nil, err
	} else {
		clog.Infof("%v order(s) listed.", len(orders))
		return &orders, nil
	}

}

func (agent *CheckoutAgent) Create(checkout *Checkout) (*Checkout, error) {
	urlStr := "/usageapi/v1/orders"
	plan, err := agent.Market.Get(checkout.PlanId)

	if err != nil {
		clog.Error(err)
		return nil, err
	}

	clog.Debugf("%+v", plan)

	if plan.PlanId == "" {
		return nil, ErrorNew(ErrCodePlanNotFound)
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := agent.BaseURL.ResolveReference(rel)

	req, err := agent.NewRequest("POST", u.String(), checkout)
	if err != nil {
		return nil, err
	}

	response := new(checkoutResponse)
	if err := agent.Do(req, response); err != nil {
		clog.Error(err)
		return nil, err
	}

	checkout.OrderID = response.OrderID

	return checkout, nil
}

type checkoutResponse struct {
	Code    uint   `json:"code"`
	Msg     string `json:"msg"`
	OrderID string `json:"data,omitempty"`
}

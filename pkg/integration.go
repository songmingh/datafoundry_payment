package pkg

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/zonesan/clog"
)

type IntegrationAgent struct {
	*Agent
	BaseURL *url.URL
}

type Integration struct {
	Amount  float64 `json:"amount,omitempty"`
	Project string  `json:"namespace,omitempty"`
}

func (agent *IntegrationAgent) ListRepos(r *http.Request) (*Integration, error) {

	urlStr := "/integration/v1/repositories"

	integration := new(Integration)

	if err := doRequest(agent, r, "GET", urlStr, nil, integration); err != nil {
		clog.Error(err)

		return nil, err
	}
	clog.Debug("...")
	return integration, nil

}

func (agent *IntegrationAgent) GetRepo(r *http.Request, repo string) (*[]PurchasedOrder, error) {

	urlStr := fmt.Sprintf("/integration/v1/repository/%v", repo)

	orders := new([]PurchasedOrder)

	if err := doRequestList(agent, r, "GET", urlStr, nil, orders); err != nil {
		clog.Error(err)
		return nil, err
	}

	clog.Infof("%v order(s) listed.", len(*orders))

	return orders, nil
}

func (agent *IntegrationAgent) ListItems(r *http.Request) (*[]PurchasedOrder, error) {
	urlStr := "/usageapi/v1/orders"

	orders := new([]PurchasedOrder)

	if err := doRequestList(agent, r, "GET", urlStr, nil, orders); err != nil {
		clog.Error(err)
		return nil, err
	}

	clog.Infof("%v order(s) listed.", len(*orders))

	return orders, nil
}

func (agent *IntegrationAgent) GetItem(r *http.Request) (*[]PurchasedOrder, error) {
	urlStr := "/usageapi/v1/orders"

	orders := new([]PurchasedOrder)

	if err := doRequestList(agent, r, "GET", urlStr, nil, orders); err != nil {
		clog.Error(err)
		return nil, err
	}

	clog.Infof("%v order(s) listed.", len(*orders))

	return orders, nil
}

func (agent *IntegrationAgent) Url() *url.URL {
	u := new(url.URL)
	u, _ = url.Parse(agent.BaseURL.String())
	return u
}

func (agent *IntegrationAgent) Instance() *Agent {
	return agent.Agent
}

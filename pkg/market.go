package pkg

import (
	"fmt"
	"net/url"

	"github.com/zonesan/clog"
)

type MarketAgent struct {
	*Agent
	BaseURL *url.URL
}

type Plan struct {
	PlanId       string  `json:"plan_id"`
	Type         string  `json:"type"`
	Price        float32 `json:"price"`
	BillPeriod   string  `json:"bill_period"`
	Desc         string  `json:"description"`
	CreationTime string  `json:"creation_time,omitempty"`
}

type Market struct {
	Plans *[]Plan `json:"plans"`
}

func (agent *MarketAgent) Get(id string) (*Plan, error) {
	// market := fakeMarket()

	urlStr := fmt.Sprintf("/charge/v1/plans/%v", id)

	plan := new(Plan)

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := agent.BaseURL.ResolveReference(rel)

	req, err := agent.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	if err := agent.Do(req, plan); err != nil {
		return nil, err
	}

	return plan, nil

}

func (agent *MarketAgent) List() (*Market, error) {
	urlStr := "/charge/v1/plans"

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := agent.BaseURL.ResolveReference(rel)

	req, err := agent.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	market := new(Market)

	response := new(marketResponse)
	if err := agent.Do(req, response); err != nil {
		clog.Error(err)
		return nil, err
	}

	return market, nil
}

type marketResponse struct {
	Code            uint   `json:"code"`
	Msg             string `json:"msg"`
	QueryListResult `json:"data,omitempty"`
}

type QueryListResult struct {
	Total   int64        `json:"total"`
	Results []marketPlan `json:"results"`
}

type marketPlan struct {
	id             int `json:"plan_id, omitempty"`
	Plan_id        string
	Plan_type      string
	Specification1 string
	Specification2 string
	Price          float32
	Cycle          string
	Create_time    string
	Status         string
}

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
	Plans []Plan `json:"plans"`
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

	for _, result := range response.Results {
		plan := Plan{
			PlanId:       result.Plan_id,
			Type:         result.Plan_type,
			Price:        result.Price,
			BillPeriod:   result.Cycle,
			Desc:         fmt.Sprintf("%v %v", result.Spec1, result.Spec2),
			CreationTime: result.Create_time,
		}
		market.Plans = append(market.Plans, plan)
	}

	return market, nil
}

type marketResponse struct {
	Code            uint   `json:"code"`
	Msg             string `json:"msg"`
	QueryListResult `json:"data,omitempty"`
}

type QueryListResult struct {
	Total   int64     `json:"total"`
	Results []apiPlan `json:"results"`
}

type apiPlan struct {
	id          int
	Plan_id     string  `json:"plan_id,omitempty"`
	Plan_type   string  `json:"type,omitempty"`
	Spec1       string  `json:"spec1,omitempty"`
	Spec2       string  `json:"spec2,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Cycle       string  `json:"bill_period,omitempty"`
	Create_time string  `json:"creation_time,omitempty"`
	Status      string  `json:"status,omitempty"`
}

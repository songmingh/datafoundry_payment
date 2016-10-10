package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	// "reflect"

	"github.com/zonesan/clog"
)

type MarketAgent struct {
	*Agent
	BaseURL *url.URL
}

type Plan struct {
	PlanId       string  `json:"plan_id"`
	Name         string  `json:"plan_name"`
	Type         string  `json:"type"`
	Price        float32 `json:"price"`
	BillPeriod   string  `json:"bill_period"`
	Desc         string  `json:"description"`
	Desc2        string  `json:"description2,omitempty"`
	Region       string  `json:"region,omitempty"`
	CreationTime string  `json:"creation_time,omitempty"`
}

type Market struct {
	Plans []Plan `json:"plans"`
}

func (agent *MarketAgent) Get(id string) (*Plan, error) {
	// market := fakeMarket()

	urlStr := fmt.Sprintf("/charge/v1/plans/%v", id)

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := agent.BaseURL.ResolveReference(rel)

	req, err := agent.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	plan := new(Plan)
	response := new(planResponse)

	if err := agent.Do(req, response); err != nil {
		return nil, err
	}

	plan.PlanId = response.Plan_id
	plan.Name = response.Name
	plan.Type = response.Plan_type
	plan.Price = response.Price
	plan.BillPeriod = response.Cycle
	plan.Region = response.Region
	plan.Desc = response.Spec1
	plan.Desc2 = response.Spec2
	plan.CreationTime = response.Create_time

	return plan, nil

}

func (agent *MarketAgent) List(r *http.Request) (*Market, error) {
	urlStr := "/charge/v1/plans"

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

	market := new(Market)

	response := new(RemoteListResponse)
	if err := agent.Do(req, response); err != nil {
		clog.Error(err)
		return nil, err
	}

	plans := []apiPlan{}

	if err := json.Unmarshal([]byte(response.Data), &plans); err != nil {
		clog.Error(err)
		return nil, err
	} else {
		for _, result := range plans {
			plan := Plan{
				PlanId:       result.Plan_id,
				Name:         result.Name,
				Type:         result.Plan_type,
				Price:        result.Price,
				BillPeriod:   result.Cycle,
				Region:       result.Region,
				Desc:         result.Spec1,
				Desc2:        result.Spec2,
				CreationTime: result.Create_time,
			}
			market.Plans = append(market.Plans, plan)
		}
	}

	return market, nil
}

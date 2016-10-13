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
	Level        int     `json:"plan_level"`
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

func (agent *MarketAgent) Get(r *http.Request, id string) (*Plan, error) {
	// market := fakeMarket()

	urlStr := fmt.Sprintf("/charge/v1/plans/%v", id)

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

	response := new(RemoteResponse)

	if err := agent.Do(req, response); err != nil {
		return nil, err
	}

	apiplan := new(apiPlan)
	plan := new(Plan)

	if err := json.Unmarshal([]byte(response.Data), apiplan); err != nil {
		clog.Error(err)
		return nil, err
	} else {
		clog.Infof("%#v", apiplan)
		plan.PlanId = apiplan.Plan_id
		plan.Name = apiplan.Name
		plan.Level = apiplan.Level
		plan.Type = apiplan.Plan_type
		plan.Price = apiplan.Price
		plan.BillPeriod = apiplan.Cycle
		plan.Region = apiplan.Region
		plan.Desc = apiplan.Spec1
		plan.Desc2 = apiplan.Spec2
		plan.CreationTime = apiplan.Create_time
	}

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
				Level:        result.Level,
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

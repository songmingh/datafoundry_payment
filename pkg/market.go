package pkg

type MarketAgent service

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

func (*MarketAgent) Get() *Market {
	market := fakeMarket()
	return market
}

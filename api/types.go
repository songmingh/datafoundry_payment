package api

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status,omitempty"`
	//Data    interface{} `json:"data,omitempty"`
}

type Account struct {
	Purchased bool   `json:"purchased"`
	Notify    bool   `json:"notification"`
	Plans     []Plan `json:"plans,omitempty"`
}

type Amount struct {
	Id           string  `json:"trans_id"`
	CreationTime string  `json:"creation_time"`
	Amount       float32 `json:"amount"`
	Desc         string  `json:"description"`
	Payment      string  `json:"payment_method"`
	Status       string  `json:"status"`
}

type Balance struct {
	Balance float32 `json:"balance"`
}

type Checkout struct {
	PlanId  string `json:"plan_id"`
	Project string `json:"namespace,omitempty"`
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

type Recharge struct {
	Amount  float32 `json:"amount"`
	Project string  `json:"namespace,omitempty"`
}

type Summary struct {
}

type Coupon struct{}

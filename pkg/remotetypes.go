package pkg

import (
	"encoding/json"
	"time"
)

type RemoteResponse struct {
	Code    uint            `json:"code"`
	Message string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
}

type RemoteListResponse struct {
	Code            uint   `json:"code"`
	Msg             string `json:"msg"`
	QueryListResult `json:"data,omitempty"`
}

type QueryListResult struct {
	Total int64           `json:"total"`
	Data  json.RawMessage `json:"results"`
}

type apiPlan struct {
	id          int
	Plan_id     string  `json:"plan_id,omitempty"`
	Name        string  `json:"plan_name,omitempty"`
	Plan_type   string  `json:"plan_type,omitempty"`
	Spec1       string  `json:"specification1,omitempty"`
	Spec2       string  `json:"specification2,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Cycle       string  `json:"cycle,omitempty"`
	Region      string  `json:"region,omitempty"`
	Create_time string  `json:"creation_time,omitempty"`
	Status      string  `json:"status,omitempty"`
}

type apiPurchaseOrder struct {
	Order_id        string     `json:"orderId,omitempty"`
	Account_id      string     `json:"project,omitempty"` // accountId
	Region          string     `json:"region,omitempty"`
	Quantities      int        `json:"quantities,omitempty"`
	Plan_id         string     `json:"planId,omitempty"`
	Plan_type       string     `json:"_,omitempty"`
	Start_time      time.Time  `json:"startTime,omitempty"`
	End_time        time.Time  `json:"_,omitempty"`       // po
	EndTime         *time.Time `json:"endTime,omitempty"` // vo
	Deadline_time   time.Time  `json:"deadline,omitempty"`
	Last_consume_id int        `json:"_,omitempty"`
	Status          int        `json:"status,omitempty"`
	Creator         string     `json:"creator,omitempty"`
}

type apiBalance struct {
	Namespace string  `json:"namespace"`
	CreateAt  string  `json:"create_at"`
	UpdateAt  string  `json:"update_at"`
	Balance   float32 `json:"balance"`
	Status    string  `json:"state,omitempty"`
}

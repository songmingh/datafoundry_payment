package pkg

import (
	"encoding/json"
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
	Level       int     `json:"plan_level,omitempty"`
	Spec1       string  `json:"specification1,omitempty"`
	Spec2       string  `json:"specification2,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Cycle       string  `json:"cycle,omitempty"`
	Region      string  `json:"region,omitempty"`
	Create_time string  `json:"creation_time,omitempty"`
	Status      string  `json:"status,omitempty"`
}

type apiPurchaseOrder struct {
	Order_id        string `json:"order_id,omitempty"`
	Account_id      string `json:"project,omitempty"` // accountId
	Region          string `json:"region,omitempty"`
	Plan_id         string `json:"plan_id,omitempty"`
	Plan_type       string `json:"_,omitempty"`
	Start_time      string `json:"start_time,omitempty"`
	End_time        string `json:"_,omitempty"`        // po
	EndTime         string `json:"end_time,omitempty"` // vo
	Deadline_time   string `json:"deadline,omitempty"`
	Last_consume_id int    `json:"_,omitempty"`
	Status          string `json:"status,omitempty"`
	Creator         string `json:"creator,omitempty"`
}

type apiBalance struct {
	Namespace string  `json:"namespace"`
	CreateAt  string  `json:"create_at"`
	UpdateAt  string  `json:"update_at"`
	Balance   float64 `json:"balance"`
	Status    string  `json:"state,omitempty"`
}

type apiTransaction struct {
	TransactionId string  `json:"transactionId"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	Namespace     string  `json:"namespace"`
	User          string  `json:"user,omitempty"`
	Reason        string  `json:"reason,omitempty"`
	CreateTime    string  `json:"createtime,omitempty"`
	Status        string  `json:"status,omitempty"`
	StatusTime    string  `json:"statustime,omitempty"`
}

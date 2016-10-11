package pkg

import (
	"net/http"
)

type AmountAgent service

type Amount struct {
	Id           string  `json:"trans_id"`
	CreationTime string  `json:"creation_time"`
	Amount       float32 `json:"amount"`
	Desc         string  `json:"description"`
	Payment      string  `json:"payment_method"`
	Comment      string  `json:"comment"`
	Status       string  `json:"status"`
}

func (u *AmountAgent) Get(r *http.Request) *Amount {
	amount := fakeAmount(r)
	return amount
}
func (u *AmountAgent) List(r *http.Request) *[]Amount {
	amounts := fakeAmounts(r)
	return amounts
}

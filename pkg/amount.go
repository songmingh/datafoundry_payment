package pkg

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/zonesan/clog"
)

type AmountAgent struct {
	*Agent
	BaseURL *url.URL
}

type Amount struct {
	Id           string  `json:"trans_id"`
	CreationTime string  `json:"creation_time"`
	Amount       float64 `json:"amount"`
	Desc         string  `json:"description"`
	Payment      string  `json:"payment_method"`
	Comment      string  `json:"comment"`
	Status       string  `json:"status"`
}

type Amounts struct {
	Amounts []Amount `json:"amounts"`
}

func (agent *AmountAgent) Get(r *http.Request, tid string) (*Amount, error) {

	urlStr := fmt.Sprintf("/charge/v1/recharge/%v", tid)

	amount := new(Amount)
	if err := doRequest(agent, r, "GET", urlStr, nil, amount); err != nil {
		clog.Error(err)
		return nil, err
	}
	return amount, nil
}

func (agent *AmountAgent) List(r *http.Request) (*Amounts, error) {

	urlStr := "/charge/v1/recharge"

	transactions := []apiTransaction{}
	amounts := new(Amounts)

	if err := doRequest(agent, r, "GET", urlStr, nil, &transactions); err != nil {
		clog.Error(err)
		return nil, err
	} else {
		for _, transaction := range transactions {
			amount := Amount{
				Id:           transaction.TransactionId,
				CreationTime: transaction.CreateTime,
				Amount:       transaction.Amount,
				Desc:         transaction.Type,
			}
			amounts.Amounts = append(amounts.Amounts, amount)
		}
	}
	return amounts, nil
}

func (agent *AmountAgent) Url() *url.URL {
	return agent.BaseURL
}

func (agent *AmountAgent) Instance() *Agent {
	return agent.Agent
}

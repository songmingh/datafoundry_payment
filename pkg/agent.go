package pkg

import (
	"net/http"
	"net/url"
	"sync"
)

const (
	defaultBaseURL = "http://localhost:7071"
)

// An Agent manages communication with the payment components API.
type Agent struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests.  Defaults to the public GitHub API, BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Agents used for talking to different parts of the payment components API.
	Recharge *RechargeAgent
	Checkout *CheckoutAgent
	Balance  *BalanceAgent
	Market   *MarketAgent
	Amount   *AmountAgent
	Account  *AccountAgent
	Coupon   *CouponAgent
}

type service struct {
	agent *Agent
}

func NewAgent(httpClient *http.Client) *Agent {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	agent := &Agent{client: httpClient, BaseURL: baseURL}

	agent.common.agent = agent
	agent.Account = (*AccountAgent)(&agent.common)
	agent.Amount = (*AmountAgent)(&agent.common)
	agent.Balance = (*BalanceAgent)(&agent.common)
	agent.Checkout = (*CheckoutAgent)(&agent.common)
	agent.Coupon = (*CouponAgent)(&agent.common)
	agent.Market = (*MarketAgent)(&agent.common)
	agent.Recharge = (*RechargeAgent)(&agent.common)

	return agent
}

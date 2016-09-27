package pkg

import (
	"github.com/asiainfoLDP/datafoundry_payment/api"
)

type MarketAgent service

func (*MarketAgent) Get() *api.Market {
	market := &api.Market{}
	return market
}

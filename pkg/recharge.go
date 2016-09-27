package pkg

import (
	"github.com/asiainfoLDP/datafoundry_payment/api"
)

type RechargeAgent service

func (*RechargeAgent) Get() *api.Recharge {
	recharge := &api.Recharge{}
	return recharge
}

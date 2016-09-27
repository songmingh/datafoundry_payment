package agent

import (
	"github.com/asiainfoLDP/datafoundry_payment/api"
)

type CheckoutAgent service

func (*CheckoutAgent) Get() *api.Checkout {
	checkout := &api.Checkout{}
	return checkout
}

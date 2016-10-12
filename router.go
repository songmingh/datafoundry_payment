package main

import (
	"github.com/asiainfoLDP/datafoundry_payment/api"
	"github.com/asiainfoLDP/datafoundry_payment/api/account"
	"github.com/asiainfoLDP/datafoundry_payment/api/amount"
	"github.com/asiainfoLDP/datafoundry_payment/api/balance"
	"github.com/asiainfoLDP/datafoundry_payment/api/checkout"
	"github.com/asiainfoLDP/datafoundry_payment/api/coupon"
	"github.com/asiainfoLDP/datafoundry_payment/api/market"
	"github.com/asiainfoLDP/datafoundry_payment/api/recharge"
	"github.com/julienschmidt/httprouter"
)

const PATHPREFIX = "/payment/v1"

func createRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", api.Index)

	router.POST(PATHPREFIX+"/recharge", recharge.Recharge)
	router.POST(PATHPREFIX+"/checkout", checkout.Checkout)
	router.GET(PATHPREFIX+"/balance", balance.Balance)
	router.GET(PATHPREFIX+"/market", market.Market)
	router.GET(PATHPREFIX+"/amounts", amount.AmountList)
	router.GET(PATHPREFIX+"/amounts/:tid", amount.Amount)
	router.GET(PATHPREFIX+"/account", account.Account)
	router.GET(PATHPREFIX+"/coupon", coupon.Coupon)

	router.NotFound = &api.Mux{}

	return router
}

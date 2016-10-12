package checkout

import (
	"net/http"

	"github.com/asiainfoLDP/datafoundry_payment/api"
	"github.com/asiainfoLDP/datafoundry_payment/pkg"
	"github.com/julienschmidt/httprouter"
	"github.com/zonesan/clog"
)

func Checkout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	checkout := new(pkg.Checkout)

	if err := api.ParseRequestBody(r, checkout); err != nil {
		clog.Error("read request body error.", err)
		api.RespError(w, err)
		return
	}

	agent := api.Agent()

	checkoutResult, err := agent.Checkout.Create(r, checkout)
	if err != nil {
		clog.Error(err)
		api.RespError(w, err)
		return
	}

	api.RespOK(w, checkoutResult)
}

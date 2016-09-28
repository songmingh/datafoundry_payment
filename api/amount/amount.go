package amount

import (
	"net/http"

	"github.com/asiainfoLDP/datafoundry_payment/api"
	"github.com/julienschmidt/httprouter"
	"github.com/zonesan/clog"
)

func AmountList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)
	agent := api.Agent()
	amounts := agent.Amount.List(r)

	api.RespOK(w, amounts)
	return
}

func Amount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	agent := api.Agent()
	amount := agent.Amount.Get(r)

	api.RespOK(w, amount)
	return
}

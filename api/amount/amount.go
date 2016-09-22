package amount

import (
	"net/http"

	"github.com/asiainfoLDP/datafoundry_payment/api"
	"github.com/asiainfoLDP/datafoundry_payment/fake"
	"github.com/julienschmidt/httprouter"
	"github.com/zonesan/clog"
)

func AmountList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	amounts := fake.Amounts(r)

	api.RespOK(w, amounts)
	return
}

func Amount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	amount := fake.Amount(r)

	api.RespOK(w, amount)
	return
}

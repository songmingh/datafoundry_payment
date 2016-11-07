package integration

import (
	"net/http"

	"github.com/asiainfoLDP/datafoundry_payment/api"
	"github.com/julienschmidt/httprouter"
	"github.com/zonesan/clog"
)

func ListRepos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	agent := api.Agent()
	if resp, err := agent.Integration.ListRepos(r); err != nil {
		api.RespError(w, err)
	} else {
		api.RespOK(w, resp)
	}
}

func GetRepo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	repo := ps.ByName("repo")

	agent := api.Agent()
	if resp, err := agent.Integration.GetRepo(r, repo); err != nil {
		api.RespError(w, err)
	} else {
		api.RespOK(w, resp)
	}
}

func ListItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	agent := api.Agent()
	if resp, err := agent.Recharge.Notification(r); err != nil {
		api.RespError(w, err)
	} else {
		api.RespOK(w, resp)
	}
}

func GetItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)

	agent := api.Agent()
	if resp, err := agent.Recharge.Notification(r); err != nil {
		api.RespError(w, err)
	} else {
		api.RespOK(w, resp)
	}
}

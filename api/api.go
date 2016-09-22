package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zonesan/clog"
)

type Mux struct{}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)
	RespError(w, ErrorNew(ErrCodeNotFound), http.StatusNotFound)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	clog.Info("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)
	fmt.Fprint(w, "Welcome!\n")
}

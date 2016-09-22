package api

import (
	// "github.com/golang/glog"
	// "github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"
)

func RespError(w http.ResponseWriter, err error, httpCode int) {
	resp := genRespJson(httpCode, err)

	if body, err := json.MarshalIndent(resp, "", "  "); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		w.Write(body)
	}

}

func RespOK(w http.ResponseWriter, data interface{}) {
	if data == nil {
		data = genRespJson(http.StatusOK, nil)
	}

	if body, err := json.MarshalIndent(data, "", "  "); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func genRespJson(httpCode int, err error) *APIResponse {
	resp := new(APIResponse)
	var msgCode int
	var message string

	if err == nil {
		msgCode = ErrCodeOK
		message = ErrText(msgCode)
	} else {
		if e, ok := err.(Error); ok {
			msgCode = e.Code
			message = e.Message
		} else if e, ok := err.(*Error); ok {
			msgCode = e.Code
			message = e.Message
		} else {
			msgCode = ErrCodeUnknownError
			message = err.Error()
		}
	}

	resp.Code = msgCode
	resp.Message = message
	resp.Status = http.StatusText(httpCode)
	return resp
}

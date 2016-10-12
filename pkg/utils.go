package pkg

import (
	"net/http"
)

func getToken(r *http.Request) (string, error) {
	r.ParseForm()

	token := r.Header.Get("Authorization")

	if token == "" {
		err := ErrorNew(ErrCodeUnauthorized)
		return "", err
	}
	return token, nil
}

func doRequest(agent AgentInterface, r *http.Request, urlStr string, reqBody, respBody interface{}) error {
	agent.BaseURL()
	return nil
}

type AgentInterface struct {
	*Agent
	Interface
}

type Interface interface {
	BaseURL() string
}

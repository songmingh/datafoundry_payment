package pkg

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/zonesan/clog"
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

type AgentInterface interface {
	Url() *url.URL
	Instance() *Agent
}

func doRequest(agent AgentInterface, r *http.Request, method, urlStr string, reqBody, respBody interface{}) error {
	baseURL := agent.Url()
	client := agent.Instance()

	if r.URL.RawQuery != "" {
		urlStr += "?" + r.URL.RawQuery
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	u := baseURL.ResolveReference(rel)

	token, err := getToken(r)
	if err != nil {
		clog.Error(err)
		return err
	}

	req, err := client.NewRequest(strings.ToUpper(method), u.String(), reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", token)

	response := new(RemoteResponse)

	if err := client.Do(req, response); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(response.Data), respBody); err != nil {
		clog.Error(err)
		return err
	}

	return nil
}

func doRequestList(agent AgentInterface, r *http.Request, method, urlStr string, reqBody, respBody interface{}) error {
	baseURL := agent.Url()
	client := agent.Instance()

	if r.URL.RawQuery != "" {
		urlStr += "?" + r.URL.RawQuery
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	u := baseURL.ResolveReference(rel)

	token, err := getToken(r)
	if err != nil {
		clog.Error(err)
		return err
	}

	req, err := client.NewRequest(strings.ToUpper(method), u.String(), reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", token)

	response := new(RemoteListResponse)

	if err := client.Do(req, response); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(response.Data), respBody); err != nil {
		clog.Error(err)
		return err
	}

	return nil
}

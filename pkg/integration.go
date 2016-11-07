package pkg

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/zonesan/clog"
)

type IntegrationAgent struct {
	*Agent
	BaseURL *url.URL
}

type Integration struct {
	Amount  float64 `json:"amount,omitempty"`
	Project string  `json:"namespace,omitempty"`
}

type DataRepo struct {
	RepoID int        `json:"repo_id,omitempty"`
	Name   string     `json:"repo_name,omitempty"`
	Class  string     `json:"class,omitempty"`
	Label  string     `json:"label,omitempty"`
	Desc   string     `json:"description,omitempty"`
	Owner  string     `json:"owner,omitempty"`
	Items  []DataItem `json:"items,omitempty"`
}

type DataItem struct {
	ItemID   int            `json:"item_id,omitempty"`
	Name     string         `json:"item_name,omitempty"`
	Url      string         `json:"url,omitempty"`
	UpdateAt string         `json:"update_at,omitempty"`
	Sameple  string         `json:"sample,omitempty"`
	Attrs    []DataItemAttr `json:"attrs,omitempty"`
}

type DataItemAttr struct {
	Name    string `json:"name,omitempty"`
	Comment string `json:"instruction,omitempty"`
	Example string `json:"example,omitempty"`
}

func (agent *IntegrationAgent) ListRepos(r *http.Request) (*[]DataRepo, error) {

	urlStr := "/integration/v1/repositories"

	apirepos := new([]DataRepo)
	repos := []DataRepo{}

	if err := doRequest(agent, r, "GET", urlStr, nil, apirepos); err != nil {
		clog.Error(err)
		return nil, err
	} else {
		for _, apirepo := range *apirepos {
			repo := DataRepo{
				RepoID: apirepo.RepoID,
				Name:   apirepo.Name,
				Class:  apirepo.Class,
				Label:  apirepo.Label,
				Desc:   apirepo.Desc,
			}
			repos = append(repos, repo)
		}
	}

	clog.Infof("%v repo(s) listed.", len(repos))
	return &repos, nil

}

func (agent *IntegrationAgent) GetRepo(r *http.Request, repoName string) (*DataRepo, error) {

	urlStr := fmt.Sprintf("/integration/v1/repository/%v", repoName)

	apirepo := new(apiDataRepo)
	repo := new(DataRepo)

	if err := doRequestList(agent, r, "GET", urlStr, nil, apirepo); err != nil {
		clog.Error(err)
		return nil, err
	}

	return repo, nil
}

func (agent *IntegrationAgent) ListItems(r *http.Request) (*[]PurchasedOrder, error) {
	urlStr := ""
	clog.Debug(urlStr)
	return nil, nil

	// orders := new([]PurchasedOrder)

	// if err := doRequestList(agent, r, "GET", urlStr, nil, orders); err != nil {
	// 	clog.Error(err)
	// 	return nil, err
	// }

	// clog.Infof("%v order(s) listed.", len(*orders))

	// return orders, nil
}

func (agent *IntegrationAgent) GetItem(r *http.Request, repo, item string) (*[]PurchasedOrder, error) {
	urlStr := fmt.Sprintf("/integration/v1/dataitem/%v/%v", repo, item)

	orders := new([]PurchasedOrder)

	if err := doRequestList(agent, r, "GET", urlStr, nil, orders); err != nil {
		clog.Error(err)
		return nil, err
	}

	clog.Infof("%v order(s) listed.", len(*orders))

	return orders, nil
}

func (agent *IntegrationAgent) Url() *url.URL {
	u := new(url.URL)
	u, _ = url.Parse(agent.BaseURL.String())
	return u
}

func (agent *IntegrationAgent) Instance() *Agent {
	return agent.Agent
}

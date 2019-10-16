package client

import (
	"fmt"
	"github.com/benschw/satis-go/satis/satisphp/api"
	"log"
	"net/http"
)

var _ = log.Print

type SatisClient struct {
	Host string
}

func (c *SatisClient) AddRepo(repo *api.Repo) (*api.Repo, error) {
	r := &api.Repo{}
	url := c.Host + "/api/repo"

	req, err := makeRequest("POST", url, repo)
	if err != nil {
		return r, err
	}
	err = processResponseEntity(req, &r, http.StatusCreated)
	return r, err
}
func (c *SatisClient) SaveRepo(repo *api.Repo) (*api.Repo, error) {
	r := &api.Repo{}
	url := c.Host + "/api/repo/" + repo.Id

	req, err := makeRequest("PUT", url, repo)
	if err != nil {
		return r, err
	}
	err = processResponseEntity(req, &r, http.StatusOK)
	return r, err
}

func (c *SatisClient) FindRepo(id string) (*api.Repo, error) {
	var repo api.Repo
	url := c.Host + "/api/repo/" + id

	req, err := makeRequest("GET", url, nil)
	if err != nil {
		return &repo, err
	}
	err = processResponseEntity(req, &repo, http.StatusOK)
	return &repo, err
}

func (c *SatisClient) FindAllRepos() ([]api.Repo, error) {
	var repos []api.Repo
	url := c.Host + "/api/repo"

	req, err := makeRequest("GET", url, nil)
	if err != nil {
		return repos, err
	}
	err = processResponseEntity(req, &repos, http.StatusOK)
	return repos, err
}

func (c *SatisClient) DeleteRepo(id string) error {
	url := c.Host + "/api/repo/" + id

	req, err := makeRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	return processResponseEntity(req, nil, http.StatusNoContent)
}

func (c *SatisClient) GenerateStaticWeb() error {
	url := c.Host + "/api/generate-web-job"

	req, err := makeRequest("POST", url, nil)
	if err != nil {
		return err
	}
	return processResponseEntity(req, nil, http.StatusCreated)
}

func (c *SatisClient) GeneratePackageStaticWeb(packageName string) error {
	url := fmt.Sprintf("%s/api/generate-package", c.Host)

	req, err := makeRequest("POST", url, &api.Package{Name: packageName})
	if err != nil {
		return err
	}
	return processResponseEntity(req, nil, http.StatusCreated)
}

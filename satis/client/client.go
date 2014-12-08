package client

import (
	"github.com/benschw/satis-go/satis/satisphp/api"
	"log"
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
	err = processResponseEntity(req, &r, 201)
	return r, err
}

func (c *SatisClient) FindAll() ([]api.Repo, error) {
	var repos []api.Repo
	url := c.Host + "/api/repo"

	req, err := makeRequest("GET", url, nil)
	if err != nil {
		return repos, err
	}
	err = processResponseEntity(req, &repos, 201)
	return repos, err
}

func (c *SatisClient) DeleteRepo(id string) error {
	url := c.Host + "/api/repo/" + id

	req, err := makeRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	return processResponseEntity(req, nil, 204)
}

func (c *SatisClient) Generate() error {
	url := c.Host + "/api/generate-web-job"

	req, err := makeRequest("POST", url, nil)
	if err != nil {
		return err
	}
	return processResponseEntity(req, nil, 201)
}

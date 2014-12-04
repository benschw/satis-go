package satisphp

import (
	"encoding/json"
	"io/ioutil"
)

type SatisDb struct {
	Name         string            `json:"name"`
	Homepage     string            `json:"homepage"`
	Repositories []SatisRepository `json:"repositories"`
	RequireAll   bool              `json:"require-all"`
}

type SatisRepository struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type SatisDbManager struct {
	Path string
	Db   SatisDb
}

func (c *SatisDbManager) Load() error {

	content, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &c.Db); err != nil {
		return err
	}
	return nil
}

func (c *SatisDbManager) Write() error {
	b, err := json.MarshalIndent(c.Db, "", "    ") // pretty print
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(c.Path, b, 0644); err != nil {
		return err
	}
	return nil
}

func (c *SatisDbManager) SaveRepo(repo SatisRepository) error {
	return nil
}

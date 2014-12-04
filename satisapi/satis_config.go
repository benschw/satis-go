package satisapi

import (
	"encoding/json"
	"io/ioutil"
)

type SatisConfig struct {
	Name         string            `json:"name"`
	Homepage     string            `json:"homepage"`
	Repositories []SatisRepository `json:"repositories"`
	RequireAll   bool              `json:"requore-all"`
}

type SatisRepository struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type SatisConfigManager struct {
	Path   string
	Config SatisConfig
}

func (c *SatisConfigManager) loadConfig() error {

	content, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &c.Config); err != nil {
		return err
	}
	return nil
}

func (c *SatisConfigManager) writeConfig() error {
	b, err := json.MarshalIndent(c.Config, "", "    ") // pretty print
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(c.Path, b, 0644); err != nil {
		return err
	}
	return nil
}

func (c *SatisConfigManager) saveRepo(repo SatisRepository) error {
	return nil
}

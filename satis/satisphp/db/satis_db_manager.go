package db

import (
	"encoding/json"
	"io/ioutil"
)

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

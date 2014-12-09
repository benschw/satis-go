package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var _ = log.Print

const (
	DbFile      = "/db.json"
	StagingFile = "/stage.json"
)

type SatisDbManager struct {
	Path string
	Db   SatisDb
}

func (c *SatisDbManager) Load() error {

	content, err := ioutil.ReadFile(c.Path + DbFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &c.Db); err != nil {
		return err
	}
	return nil
}

func (c *SatisDbManager) Write() error {
	return c.doWrite(c.Path + DbFile)
}
func (c *SatisDbManager) WriteStaging() error {
	return c.doWrite(c.Path + StagingFile)
}

func (c *SatisDbManager) doWrite(path string) error {
	b, err := json.MarshalIndent(c.Db, "", "    ") // pretty print
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, b, 0644); err != nil {
		return err
	}
	return nil
}

func (c *SatisDbManager) SaveRepo(repo SatisRepository) error {
	return nil
}

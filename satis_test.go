package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func loadTestConfig() (string, Config) {
	testPath := "./test-config.json"
	path := "./config.json.tpl"
	r := Repository{Path: path}

	cfg, _ := r.loadConfig()

	r.Path = testPath
	r.writeConfig(cfg)

	return testPath, cfg
}

func TestLoadConfig(t *testing.T) {

	// given
	path, testConfig := loadTestConfig()
	r := Repository{Path: path}

	// when
	cfg, err := r.loadConfig()

	// then
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cfg, testConfig) {
		t.Error("loaded config doesn't match original")
	}
}
func TestWriteConfig(t *testing.T) {
	// given
	path, _ := loadTestConfig()
	r := Repository{Path: path}
	testConfig, _ := r.loadConfig()

	// when
	testConfig.Name = "foo"

	err := r.writeConfig(testConfig)

	// then
	if err != nil {
		t.Error(err)
	}

	cfg, err := r.loadConfig()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cfg, testConfig) {
		t.Error("config didn't persist changes when written")
	}
}

package satisapi

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func loadTestConfig() (string, SatisConfig) {
	testPath := "../test-config.json"
	path := "../config.json.tpl"
	mgr := SatisConfigManager{Path: path}

	mgr.loadConfig()

	mgr.Path = testPath
	mgr.writeConfig()

	return testPath, mgr.Config
}

func TestLoadConfig(t *testing.T) {

	// given
	path, testConfig := loadTestConfig()
	r := SatisConfigManager{Path: path}

	// when
	err := r.loadConfig()

	// then
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(r.Config, testConfig) {
		t.Error("loaded config doesn't match original")
	}
}
func TestWriteConfig(t *testing.T) {
	// given
	path, _ := loadTestConfig()
	r := SatisConfigManager{Path: path}
	r.loadConfig()
	oldConfig := r.Config
	// when
	r.Config.Name = "foo"
	modifiedConfig := r.Config

	err := r.writeConfig()

	// then
	if err != nil {
		t.Error(err)
	}

	err = r.loadConfig()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(r.Config, oldConfig) {
		t.Error("config should have changed")
	}
	if !reflect.DeepEqual(r.Config, modifiedConfig) {
		t.Error("config didn't persist changes when written")
	}
}

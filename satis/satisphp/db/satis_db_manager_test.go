package db

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var _ = fmt.Print
var _ = log.Print

func ARandomDbMgr() SatisDbManager {
	testPath := "../../../test-db.json"
	mgr := SatisDbManager{Path: testPath}
	mgr.Db.Name = "My Repo"
	mgr.Db.Homepage = "http://repo.com"
	mgr.Db.RequireAll = true
	mgr.Db.Repositories = []SatisRepository{
		SatisRepository{Type: "vcs", Url: "http://package.com"},
	}

	mgr.Path = testPath
	mgr.Write()

	return mgr
}

func TestDbLoad(t *testing.T) {

	// given
	mgr := ARandomDbMgr()
	r := SatisDbManager{Path: mgr.Path}

	// when
	err := r.Load()

	// then
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(r.Db, mgr.Db) {
		t.Error("loaded config doesn't match original")
	}
}
func TestDbWrite(t *testing.T) {
	// given
	r := ARandomDbMgr()
	oldName := r.Db.Name
	// when
	r.Db.Name = "foo"
	modifiedDb := r.Db

	err := r.Write()

	// then
	if err != nil {
		t.Error(err)
	}

	err = r.Load()
	if err != nil {
		t.Error(err)
	}

	if oldName == r.Db.Name {
		t.Errorf("config should have changed: %s / %s", oldName, r.Db.Name)
	}
	if !reflect.DeepEqual(r.Db, modifiedDb) {
		t.Errorf("config didn't persist changes when written: %s / %s", r.Db.Name, modifiedDb.Name)
	}
}

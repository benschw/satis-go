package satisphp

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var _ = fmt.Print
var _ = log.Print

type StubGenerator struct {
	runs int
}

func (s *StubGenerator) Generate() error {
	s.runs++
	return nil
}

var gen *StubGenerator

func ARandomClient() *SatisClient {
	path := "../../test-db.json"

	dbMgr := &SatisDbManager{Path: path}
	dbMgr.Write() // empty

	jobs := make(chan SatisJob)

	// Job Processor responsible for interacting with db & static web docs
	gen = &StubGenerator{}

	jobProcessor := SatisJobProcessor{
		Jobs:      jobs,
		Generator: gen,
	}

	// Client to Job Processor
	satisClient := &SatisClient{
		DbPath: path,
		Jobs:   jobs,
	}

	// Start update processor
	go jobProcessor.ProcessUpdates()

	return satisClient
}

func TestSave(t *testing.T) {

	// given
	c := ARandomClient()
	repo := SatisRepository{
		Type: "vcs",
		Url:  "http://foo.bar",
	}

	// when
	err := c.SaveRepo(repo)

	// then
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {

	// given
	c := ARandomClient()
	repo1 := SatisRepository{
		Type: "vcs",
		Url:  "http://foo.bar",
	}
	repo2 := SatisRepository{
		Type: "vcs",
		Url:  "http://baz.boo",
	}
	c.SaveRepo(repo1)
	c.SaveRepo(repo2)

	// when
	err := c.DeleteRepo(repo1.Url)

	// then
	if err != nil {
		t.Error(err)
	}
	repos, _ := c.FindAllRepos()

	if !reflect.DeepEqual(repos, []SatisRepository{repo2}) {
		t.Errorf("repos don't match expected: %v", repos)
	}
}

func TestFindAll(t *testing.T) {

	// given
	c := ARandomClient()
	repo := SatisRepository{
		Type: "vcs",
		Url:  "http://foo.bar",
	}
	c.SaveRepo(repo)
	// when
	repos, err := c.FindAllRepos()

	// then
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual([]SatisRepository{repo}, repos) {
		t.Errorf("repos don't match expected: %v", repos)
	}

}

func TestGenerate(t *testing.T) {

	// given
	c := ARandomClient()

	// when
	err := c.GenerateSatisWeb()

	// then
	if err != nil {
		t.Error(err)
	}

	if gen.runs != 1 {
		t.Errorf("generator run wrong number of times: %d", gen.runs)
	}

}

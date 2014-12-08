package satisphp

import (
	"fmt"
	"github.com/benschw/satis-go/satis/satisphp/api"
	"github.com/benschw/satis-go/satis/satisphp/db"
	"github.com/benschw/satis-go/satis/satisphp/job"
	"log"
	"reflect"
	"testing"
	"time"
)

var _ = fmt.Print
var _ = log.Print

type StubGenerator struct {
	runs int
}

func (s *StubGenerator) Generate() error {
	time.Sleep(100 * time.Millisecond)
	s.runs++
	return nil
}

var gen *StubGenerator

func ARandomClient() *SatisClient {
	path := "../../test-db.json"

	dbMgr := &db.SatisDbManager{Path: path}
	dbMgr.Write() // empty

	jobs := make(chan job.SatisJob)

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
	repo := api.NewRepo("vcs", "http://foo.bar")

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
	repo1 := api.NewRepo("vcs", "http://foo.bar")
	repo2 := api.NewRepo("vcs", "http://baz.boo")
	c.SaveRepo(repo1)
	c.SaveRepo(repo2)

	// when
	err := c.DeleteRepo(repo1.Id)

	// then
	if err != nil {
		t.Error(err)
	}
	repos, _ := c.FindAllRepos()

	expected := []api.Repo{*repo2}
	if !reflect.DeepEqual(expected, repos) {
		t.Errorf("repos don't match expected: %v / %v", repos, expected)
	}
}

func TestFindAll(t *testing.T) {

	// given
	c := ARandomClient()
	repo := api.NewRepo("vcs", "http://foo.bar")
	c.SaveRepo(repo)
	// when
	repos, err := c.FindAllRepos()

	// then
	if err != nil {
		t.Error(err)
	}
	expected := []api.Repo{*repo}
	if !reflect.DeepEqual(expected, repos) {
		t.Errorf("repos don't match expected: %v / %v", repos, expected)
	}

}

func TestGenerate(t *testing.T) {

	// given
	c := ARandomClient()

	// when
	err := c.GenerateSatisWeb()

	// then
	c.Shutdown()
	if err != nil {
		t.Error(err)
	}

	if gen.runs != 1 {
		t.Errorf("generator run wrong number of times: %d", gen.runs)
	}
}

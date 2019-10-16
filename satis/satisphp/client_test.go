package satisphp

import (
	"fmt"
	"github.com/benschw/satis-go/satis/satisphp/api"
	"github.com/benschw/satis-go/satis/satisphp/db"
	"github.com/benschw/satis-go/satis/satisphp/job"
	"log"
	"os"
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

func (s *StubGenerator) GenerateRepo(repoPackage string) error {
	time.Sleep(100 * time.Millisecond)
	s.runs++
	return nil
}

var gen *StubGenerator

func ARandomClient() *SatisClient {
	dbPath := "/tmp/satis-test-data"

	// Make Data Dir
	if err := os.MkdirAll(dbPath, 0744); err != nil {
		log.Fatalf("Unable to create path: %v", err)
	}

	dbMgr := &db.SatisDbManager{Path: dbPath}
	dbMgr.Write() // empty

	jobs := make(chan job.SatisJob)

	// Job Processor responsible for interacting with db & static web docs
	gen = &StubGenerator{}

	jobProcessor := SatisJobProcessor{
		DbPath:    dbPath,
		Jobs:      jobs,
		Generator: gen,
	}

	// Client to Job Processor
	satisClient := &SatisClient{
		DbPath: dbPath,
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
	err := c.SaveRepo(repo, false)

	// then
	if err != nil {
		t.Error(err)
	}

	c.Shutdown()
	if gen.runs != 0 {
		t.Errorf("generator run wrong number of times: %d", gen.runs)
	}
}

func TestSaveAndGenerate(t *testing.T) {

	// given
	c := ARandomClient()
	repo := api.NewRepo("vcs", "http://foo.bar")

	// when
	err := c.SaveRepo(repo, true)

	// then
	if err != nil {
		t.Error(err)
	}

	c.Shutdown()
	if gen.runs != 1 {
		t.Errorf("generator run wrong number of times: %d", gen.runs)
	}
}

func TestFindAll(t *testing.T) {

	// given
	c := ARandomClient()
	repo := api.NewRepo("vcs", "http://foo.bar")
	c.SaveRepo(repo, false)
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

func TestDelete(t *testing.T) {

	// given
	c := ARandomClient()
	repo1 := api.NewRepo("vcs", "http://foo.bar")
	repo2 := api.NewRepo("vcs", "http://baz.boo")
	c.SaveRepo(repo1, false)
	c.SaveRepo(repo2, false)

	// when
	err := c.DeleteRepo(repo1.Id, false)

	// then
	if err != nil {
		t.Error(err)
	}
	repos, _ := c.FindAllRepos()

	expected := []api.Repo{*repo2}
	if !reflect.DeepEqual(expected, repos) {
		t.Errorf("repos don't match expected: %v / %v", repos, expected)
	}

	c.Shutdown()
	if gen.runs != 0 {
		t.Errorf("generator run wrong number of times: %d", gen.runs)
	}
}

func TestDeleteAndGenerate(t *testing.T) {

	// given
	c := ARandomClient()
	repo1 := api.NewRepo("vcs", "http://foo.bar")
	repo2 := api.NewRepo("vcs", "http://baz.boo")
	c.SaveRepo(repo1, false)
	c.SaveRepo(repo2, false)

	// when
	err := c.DeleteRepo(repo1.Id, true)

	// then
	if err != nil {
		t.Error(err)
	}
	repos, _ := c.FindAllRepos()

	expected := []api.Repo{*repo2}
	if !reflect.DeepEqual(expected, repos) {
		t.Errorf("repos don't match expected: %v / %v", repos, expected)
	}

	c.Shutdown()
	if gen.runs != 1 {
		t.Errorf("generator run wrong number of times: %d", gen.runs)
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

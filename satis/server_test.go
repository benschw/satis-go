package satis

import (
	"fmt"
	"github.com/gorilla/http"
	"log"
	"testing"
	"time"
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

var stubGenerator = &StubGenerator{}

func ARandomServer() *Server {
	s := &Server{
		DbPath:    "../test-db.json",
		WebPath:   "../test-web/",
		SatisPath: "../lib/satis/",
		Bind:      ":9090",
		Name:      "My Repo",
		Homepage:  "http://localhost:9090",
	}

	go s.Run()
	time.Sleep(100 * time.Millisecond)

	s.jobProcessor.Generator = stubGenerator
	return s
}

func TestGenerate(t *testing.T) {

	// given
	s := ARandomServer()

	// when
	status, _, r, err := http.DefaultClient.Post(s.Homepage+"/api/generate-web-job", nil, nil)

	// then
	if err != nil {
		t.Error(err)
	}
	if r != nil {
		defer r.Close()
	}
	if status.Code != 201 {
		t.Errorf("Bad Status: %v", status)
	}
	if stubGenerator.runs != 1 {
		t.Errorf("generator run wrong number of times: %d", stubGenerator.runs)
	}
}

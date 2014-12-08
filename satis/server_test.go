package satis

import (
	"fmt"
	"github.com/benschw/satis-go/satis/client"
	"github.com/benschw/satis-go/satis/satisphp/api"
	"github.com/benschw/satis-go/satis/satisphp/db"
	. "gopkg.in/check.v1"
	"log"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

var _ = fmt.Print
var _ = log.Print

func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	s             *Server
	stubGenerator *StubGenerator
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpSuite(c *C) {
	s.s = ARandomServer()
	s.stubGenerator = &StubGenerator{}
	s.s.jobProcessor.Generator = s.stubGenerator

}
func (s *MySuite) SetUpTest(c *C) {
	s.stubGenerator.runs = 0

	dbMgr := &db.SatisDbManager{Path: s.s.DbPath}
	dbMgr.Write()
}

func (s *MySuite) TestFindAll(c *C) {
	// given
	client := &client.SatisClient{Host: s.s.Homepage}
	repo := api.NewRepo("vcs", "http://foo.bar")
	created, _ := client.AddRepo(repo)

	// when
	found, err := client.FindAll()

	// then
	c.Assert(err, Equals, nil)

	c.Assert([]api.Repo{*created}, DeepEquals, found)
}
func (s *MySuite) TestAddRepo(c *C) {
	// given
	client := &client.SatisClient{Host: s.s.Homepage}
	repo := api.NewRepo("vcs", "http://foo.bar")

	// when
	created, err := client.AddRepo(repo)

	// then
	c.Assert(err, Equals, nil)

	found, _ := client.FindAll()
	c.Assert([]api.Repo{*created}, DeepEquals, found)
}

func (s *MySuite) TestDeleteRepo(c *C) {
	// given
	client := &client.SatisClient{Host: s.s.Homepage}
	repo := api.NewRepo("vcs", "http://foo.bar")
	created, _ := client.AddRepo(repo)

	// when
	err := client.DeleteRepo(created.Id)

	// then
	c.Assert(err, Equals, nil)

	found, _ := client.FindAll()
	c.Assert([]api.Repo{}, DeepEquals, found)
}

func (s *MySuite) TestGenerate(c *C) {
	// given
	client := &client.SatisClient{Host: s.s.Homepage}

	// when
	err := client.Generate()

	// then
	c.Assert(err, Equals, nil)

	c.Assert(s.stubGenerator.runs, Equals, 1)
}

// Stub Generator that doesn't require a system call
type StubGenerator struct {
	runs int
}

func (s *StubGenerator) Generate() error {
	s.runs++
	return nil
}

func ARandomServer() *Server {
	host := fmt.Sprintf("localhost:%d", GetRandomPort())

	s := &Server{
		DbPath:    "../test-db.json",
		WebPath:   "../test-web/",
		SatisPath: "../lib/satis/",
		Bind:      host,
		Name:      "My Repo",
		Homepage:  fmt.Sprintf("http://%s", host),
	}

	go s.Run()
	time.Sleep(100 * time.Millisecond)

	return s
}
func GetRandomPort() int {
	l, _ := net.Listen("tcp", ":0")
	defer l.Close()
	addrParts := strings.Split(l.Addr().String(), ":")
	port, _ := strconv.Atoi(addrParts[len(addrParts)-1])
	return port
}

package satis

import (
	"github.com/benschw/satis-go/satis/satisphp"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var _ = log.Printf

type Server struct {
	DbPath       string
	WebPath      string
	SatisPath    string
	Bind         string
	Name         string
	Homepage     string
	jobProcessor satisphp.SatisJobProcessor
}

func (s *Server) Run() error {
	// sync config to db
	if err := s.initDb(); err != nil {
		return err
	}

	// Shared Jobs Channel to queue/process db modifications and generation task
	jobs := make(chan satisphp.SatisJob)

	// Job Processor responsible for interacting with db & static web docs
	gen := &satisphp.StaticWebGenerator{
		DbPath:    s.DbPath,
		SatisPath: s.SatisPath,
		WebPath:   s.WebPath,
	}

	s.jobProcessor = satisphp.SatisJobProcessor{
		Jobs:      jobs,
		Generator: gen,
	}

	// Client to Job Processor
	satisClient := satisphp.SatisClient{
		DbPath: s.DbPath,
		Jobs:   jobs,
	}

	// route handlers
	resource := &SatisResource{
		SatisPhpClient: satisClient,
	}

	// Configure Routes
	r := mux.NewRouter()

	r.HandleFunc("/api/generate-web-job", resource.generateStaticWeb).Methods("POST")
	r.HandleFunc("/api/repo", resource.saveRepo).Methods("POST")
	r.HandleFunc("/api/repo/{id}", resource.deleteRepo).Methods("DELETE")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(s.WebPath)))

	http.Handle("/", r)

	// Start update processor
	go s.jobProcessor.ProcessUpdates()

	// Start HTTP Server
	return http.ListenAndServe(s.Bind, nil)
}

// Sync configured values to satis repository meta data
func (s *Server) initDb() error {
	dbMgr := &satisphp.SatisDbManager{Path: s.DbPath}

	// create empty db if it doesn't exist
	if _, err := os.Stat(s.DbPath); os.IsNotExist(err) {
		if err := dbMgr.Write(); err != nil {
			return err
		}
	}

	if err := dbMgr.Load(); err != nil {
		return err
	}
	dbMgr.Db.Name = s.Name
	dbMgr.Db.Homepage = s.Homepage
	dbMgr.Db.RequireAll = true
	return dbMgr.Write()
}

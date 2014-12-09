package satis

import (
	"github.com/benschw/satis-go/satis/satisphp"
	"github.com/benschw/satis-go/satis/satisphp/db"
	"github.com/benschw/satis-go/satis/satisphp/job"
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
	jobClient    satisphp.SatisClient
}

func (s *Server) Run() error {
	// sync config to db
	if err := s.initDb(); err != nil {
		return err
	}

	// Shared Jobs Channel to queue/process db modifications and generation task
	jobs := make(chan job.SatisJob)

	// Job Processor responsible for interacting with db & static web docs
	gen := &satisphp.StaticWebGenerator{
		DbPath:    s.DbPath,
		SatisPath: s.SatisPath,
		WebPath:   s.WebPath,
	}

	s.jobProcessor = satisphp.SatisJobProcessor{
		DbPath:    s.DbPath,
		Jobs:      jobs,
		Generator: gen,
	}

	// Client to Job Processor
	jobClient := satisphp.SatisClient{
		DbPath: s.DbPath,
		Jobs:   jobs,
	}

	// route handlers
	resource := &SatisResource{
		Host:           s.Homepage,
		SatisPhpClient: jobClient,
	}

	// Configure Routes
	r := mux.NewRouter()

	r.HandleFunc("/api/repo", resource.addRepo).Methods("POST")
	r.HandleFunc("/api/repo/{id}", resource.saveRepo).Methods("PUT")
	r.HandleFunc("/api/repo/{id}", resource.findRepo).Methods("GET")
	r.HandleFunc("/api/repo", resource.findAllRepos).Methods("GET")
	r.HandleFunc("/api/repo/{id}", resource.deleteRepo).Methods("DELETE")
	r.HandleFunc("/api/generate-web-job", resource.generateStaticWeb).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(s.WebPath)))

	//	r.Handle("/dist/{rest}", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))
	// r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))

	http.Handle("/", r)
	http.Handle("/admin/", http.StripPrefix("/admin/", http.FileServer(http.Dir("./admin-ui/"))))

	// Start update processor
	go s.jobProcessor.ProcessUpdates()

	// Start HTTP Server
	return http.ListenAndServe(s.Bind, nil)
}

// Sync configured values to satis repository meta data
func (s *Server) initDb() error {
	dbMgr := &db.SatisDbManager{Path: s.DbPath}

	// create empty db if it doesn't exist
	if _, err := os.Stat(s.DbPath + db.DbFile); os.IsNotExist(err) {
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

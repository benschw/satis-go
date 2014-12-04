package satisapi

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct{}

func (s *Server) Run() error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	configPath := dir + "/config.json.tpl"
	webPath := dir + "/web/"

	updateJobs := make(chan UpdateJob)

	updateProcessor := SatisUpdateProcessor{ConfigPath: configPath, UpdateJobs: updateJobs}

	go updateProcessor.ProcessUpdates()

	satisMgr := SatisManager{UpdateJobs: updateJobs}

	// set up http server
	resource := &SatisResource{
		SatisMgr:   satisMgr,
		ConfigPath: configPath,
		WebPath:    webPath,
	}
	r := mux.NewRouter()
	r.HandleFunc("/api/generate-job", resource.generateStaticWeb).Methods("POST")
	r.HandleFunc("/api/config/repo", resource.saveRepoInConfig).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(webPath)))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

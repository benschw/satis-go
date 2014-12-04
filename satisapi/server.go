package satisapi

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Server struct {
	satisMgr SatisManager
}

func (s *Server) generate(w http.ResponseWriter, r *http.Request) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	cmd := dir + "/build-web.sh"
	out, err := exec.Command(cmd).Output()

	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Problem Generating Satis Repository\n%s\n%s", err, string(out[:]))
		return
	}

	log.Print(string(out[:]))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func (s *Server) Run() error {

	updateJobs := make(chan UpdateJob)

	updateProcessor := SatisUpdateProcessor{ConfigPath: "./config.json.tpl", UpdateJobs: updateJobs}

	go updateProcessor.ProcessUpdates()

	s.satisMgr = SatisManager{UpdateJobs: updateJobs}

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/generate", s.generate)
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

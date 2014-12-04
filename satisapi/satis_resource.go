package satisapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type SatisResource struct {
	SatisMgr   SatisManager
	ConfigPath string
	WebPath    string
}

func (r *SatisResource) generateStaticWeb(res http.ResponseWriter, req *http.Request) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.
		Command(dir+"/satis/bin/satis", "build", r.ConfigPath, r.WebPath).
		Output()

	if err != nil {
		log.Print(err)

		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "Problem Generating Satis Repository\n%s\n%s", err, string(out[:]))
		return
	}

	log.Print(string(out[:]))

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "OK")
}

func (r *SatisResource) saveRepoInConfig(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	repo := &SatisRepository{}
	if err := decoder.Decode(repo); err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := r.SatisMgr.saveRepo(*repo); err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "OK")
}

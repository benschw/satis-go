package satis

import (
	"encoding/json"
	"fmt"
	"github.com/benschw/satisapi-go/satis/satisphp"
	"log"
	"net/http"
)

type SatisResource struct {
	SatisPhpClient satisphp.SatisClient
}

// Regenerate static web docs
func (r *SatisResource) generateStaticWeb(res http.ResponseWriter, req *http.Request) {

	// regenerate satis-web
	if err := r.SatisPhpClient.GenerateSatisWeb(); err != nil {
		log.Print(err)

		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "Problem Generating Satis Repository\n%s", err)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "OK")
}

// Add or update repository in Satis Repo and regenerate static web docs
func (r *SatisResource) saveRepo(res http.ResponseWriter, req *http.Request) {
	repo := &satisphp.SatisRepository{}

	// unmarshal post body
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(repo); err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// save config and regenerate satis-web
	if err := r.SatisPhpClient.SaveRepo(*repo); err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "OK")
}

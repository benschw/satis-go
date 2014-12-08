package satis

import (
	"encoding/json"
	"fmt"
	"github.com/benschw/satis-go/satis/satisphp"
	"github.com/benschw/satis-go/satis/satisphp/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type SatisResource struct {
	SatisPhpClient satisphp.SatisClient
}

// Regenerate static web docs
func (r *SatisResource) generateStaticWeb(res http.ResponseWriter, req *http.Request) {
	if err := r.SatisPhpClient.GenerateSatisWeb(); err != nil {
		log.Print(err)

		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
}

// Get All Repos
func (r *SatisResource) findAllRepos(res http.ResponseWriter, req *http.Request) {

	repos, err := r.SatisPhpClient.FindAllRepos()
	if err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// marshal response
	b, err := json.Marshal(repos)
	if err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(b[:]))
}

// Add or update repository in Satis Repo and regenerate static web docs
func (r *SatisResource) saveRepo(res http.ResponseWriter, req *http.Request) {
	repo := &api.Repo{}

	// unmarshal post body
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(repo); err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// save config and regenerate satis-web
	if err := r.SatisPhpClient.SaveRepo(repo); err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// marshal response
	newRepo := api.NewRepo(repo.Type, repo.Url)
	b, err := json.Marshal(newRepo)
	if err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Location", fmt.Sprintf("/api/repo/%d", newRepo.Id))
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(b[:]))
}

func (r *SatisResource) deleteRepo(res http.ResponseWriter, req *http.Request) {
	repoId := mux.Vars(req)["id"]

	if err := r.SatisPhpClient.DeleteRepo(repoId); err != nil {
		switch err {
		case satisphp.ErrRepoNotFound:
			res.WriteHeader(http.StatusNotFound)
		default:
			log.Print(err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNoContent)
}

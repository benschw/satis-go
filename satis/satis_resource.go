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
	Host           string
	SatisPhpClient satisphp.SatisClient
}

// Add repository in Satis Repo and regenerate static web docs
func (r *SatisResource) addRepo(res http.ResponseWriter, req *http.Request) {
	// unmarshal post body
	apiR := &api.Repo{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(apiR); err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	repo := api.NewRepo(apiR.Type, apiR.Url)

	if _, err := r.SatisPhpClient.FindRepo(repo.Id); err == nil || err != satisphp.ErrRepoNotFound {
		res.WriteHeader(http.StatusConflict)
		return
	}

	body, err := r.upsertRepo(repo)
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.Header().Set("Location", fmt.Sprintf("%s/api/repo/%d", r.Host, repo.Id))
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, body)
}

// Add repository in Satis Repo and regenerate static web docs
func (r *SatisResource) saveRepo(res http.ResponseWriter, req *http.Request) {
	repoId := mux.Vars(req)["id"]

	repo := &api.Repo{}

	// unmarshal post body
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(repo); err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if repo.Id != "" && repo.Id != repoId {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := r.SatisPhpClient.FindRepo(repoId); err != nil {
		switch err {
		case satisphp.ErrRepoNotFound:
			res.WriteHeader(http.StatusNotFound)
		default:
			log.Print(err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	body, err := r.upsertRepo(repo)
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, body)
}

// save config and regenerate satis-web
func (r *SatisResource) upsertRepo(repo *api.Repo) (string, error) {
	if err := r.SatisPhpClient.SaveRepo(repo); err != nil {
		return "", err
	}

	// marshal response
	newRepo := api.NewRepo(repo.Type, repo.Url)
	b, err := json.Marshal(newRepo)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

// Get One Repo
func (r *SatisResource) findRepo(res http.ResponseWriter, req *http.Request) {
	repoId := mux.Vars(req)["id"]

	repo, err := r.SatisPhpClient.FindRepo(repoId)

	if err != nil {
		switch err {
		case satisphp.ErrRepoNotFound:
			res.WriteHeader(http.StatusNotFound)
		default:
			log.Print(err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// marshal response
	b, err := json.Marshal(repo)
	if err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, string(b[:]))
}

// Get All Repos
func (r *SatisResource) findAllRepos(res http.ResponseWriter, req *http.Request) {

	repos, err := r.SatisPhpClient.FindAllRepos()
	if err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
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
	res.WriteHeader(http.StatusOK)
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

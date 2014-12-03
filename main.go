package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	Name         string             `json:"name"`
	Homepage     string             `json:"homepage"`
	Repositories []RepositoryConfig `json:"repositories"`
	RequireAll   bool               `json:"requore-all"`
}

type RepositoryConfig struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Repository struct {
	Path string
}

func (r *Repository) loadConfig() (Config, error) {
	var cfg Config

	content, err := ioutil.ReadFile(r.Path)
	if err != nil {
		return cfg, err
	}

	if err = json.Unmarshal(content, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (r *Repository) writeConfig(cfg Config) error {
	b, err := json.MarshalIndent(cfg, "", "    ") // pretty print
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(r.Path, b, 0644); err != nil {
		return err
	}
	return nil
}

func generate(w http.ResponseWriter, r *http.Request) {

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

func main() {

	// fs := http.FileServer(http.Dir("./web"))
	// http.Handle("/", fs)

	// http.HandleFunc("/generate", generate)
	// log.Fatal(http.ListenAndServe(":8000", nil))
}

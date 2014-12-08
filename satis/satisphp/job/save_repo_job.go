package job

import (
	"github.com/benschw/satis-go/satis/satisphp/db"
)

// Add or save a repo tp the repo collection
func NewSaveRepoJob(dbPath string, repo db.SatisRepository, gen bool) *SaveRepoJob {
	return &SaveRepoJob{
		dbPath:     dbPath,
		generate:   gen,
		repository: repo,
		exitChan:   make(chan error, 1),
	}
}

type SaveRepoJob struct {
	dbPath     string
	repository db.SatisRepository
	generate   bool
	exitChan   chan error
}

func (j SaveRepoJob) Generate() bool {
	return j.generate
}
func (j SaveRepoJob) ExitChan() chan error {
	return j.exitChan
}
func (j SaveRepoJob) Run() error {
	dbMgr := db.SatisDbManager{Path: j.dbPath}

	if err := dbMgr.Load(); err != nil {
		return err
	}
	repos, err := j.doSave(j.repository, dbMgr.Db.Repositories)
	if err != nil {
		return err
	}
	dbMgr.Db.Repositories = repos

	if err := dbMgr.Write(); err != nil {
		return err
	}
	return nil
}
func (j SaveRepoJob) doSave(repo db.SatisRepository, repos []db.SatisRepository) ([]db.SatisRepository, error) {
	found := false
	for i, r := range repos {
		if r.Url == repo.Url {
			repos[i] = repo
			found = true
		}
	}
	if !found {
		return append(repos, repo), nil
	}

	return repos, nil
}

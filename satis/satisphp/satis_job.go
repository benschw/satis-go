package satisphp

import (
	"errors"
)

const NotFoundError = "Not Found"

type SatisJob interface {
	Generate() bool
	ExitChan() chan error
	Run() error
}

func NewGenerateJob() *GenerateJob {
	return &GenerateJob{
		exitChan: make(chan error, 1),
	}
}

// No op job to signal job processor to regenerate satic web
type GenerateJob struct {
	exitChan chan error
}

func (j *GenerateJob) Generate() bool {
	return true
}
func (j *GenerateJob) ExitChan() chan error {
	return j.exitChan
}
func (j *GenerateJob) Run() error {
	return nil
}

// Add or save a repo tp the repo collection
func NewSaveRepoJob(dbPath string, repo SatisRepository, gen bool) *SaveRepoJob {
	return &SaveRepoJob{
		dbPath:     dbPath,
		generate:   gen,
		repository: repo,
		exitChan:   make(chan error, 1),
	}
}

type SaveRepoJob struct {
	dbPath     string
	repository SatisRepository
	generate   bool
	exitChan   chan error
}

func (j *SaveRepoJob) Generate() bool {
	return j.generate
}
func (j *SaveRepoJob) ExitChan() chan error {
	return j.exitChan
}
func (j *SaveRepoJob) Run() error {
	dbMgr := SatisDbManager{Path: j.dbPath}

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
func (j *SaveRepoJob) doSave(repo SatisRepository, repos []SatisRepository) ([]SatisRepository, error) {
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

// Remove a repo from the repo collection
func NewDeleteRepoJob(dbPath string, repo string, gen bool) *DeleteRepoJob {
	return &DeleteRepoJob{
		dbPath:     dbPath,
		generate:   gen,
		repository: repo,
		exitChan:   make(chan error, 1),
	}
}

type DeleteRepoJob struct {
	dbPath     string
	repository string
	generate   bool
	exitChan   chan error
}

func (j *DeleteRepoJob) Generate() bool {
	return j.generate
}
func (j *DeleteRepoJob) ExitChan() chan error {
	return j.exitChan
}
func (j *DeleteRepoJob) Run() error {
	dbMgr := SatisDbManager{Path: j.dbPath}

	if err := dbMgr.Load(); err != nil {
		return err
	}
	repos, err := j.doDelete(j.repository, dbMgr.Db.Repositories)
	if err != nil {
		return err
	}
	dbMgr.Db.Repositories = repos

	if err := dbMgr.Write(); err != nil {
		return err
	}
	return nil
}
func (j *DeleteRepoJob) doDelete(repo string, repos []SatisRepository) ([]SatisRepository, error) {
	var err error = nil
	found := false

	rs := make([]SatisRepository, 0, len(repos))
	for _, r := range repos {
		if r.Url == repo {
			found = true
		} else {
			rs = append(rs, r)
		}
	}
	if !found {
		err = errors.New(NotFoundError)
	}
	return rs, err
}

//Get all repos in the repo collection
func NewFindAllJob(dbPath string) *FindAllJob {
	return &FindAllJob{
		dbPath:    dbPath,
		exitChan:  make(chan error, 1),
		reposResp: make(chan []SatisRepository, 1),
	}
}

type FindAllJob struct {
	dbPath    string
	reposResp chan []SatisRepository
	generate  bool
	exitChan  chan error
}

func (j *FindAllJob) Generate() bool {
	return j.generate
}
func (j *FindAllJob) ExitChan() chan error {
	return j.exitChan
}
func (j *FindAllJob) Run() error {
	dbMgr := SatisDbManager{Path: j.dbPath}

	err := dbMgr.Load()

	j.reposResp <- dbMgr.Db.Repositories // might be empty

	return err // might be nil
}

// No op Job to signal job processor to exit
func NewExitJob() *ExitJob {
	return &ExitJob{
		exitChan: make(chan error, 1),
	}
}

type ExitJob struct {
	exitChan chan error
}

func (j *ExitJob) Generate() bool {
	return false
}
func (j *ExitJob) ExitChan() chan error {
	return j.exitChan
}
func (j *ExitJob) Run() error {
	return nil
}

package job

import (
	"github.com/benschw/satis-go/satis/satisphp/db"
)

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

func (j DeleteRepoJob) Generate() bool {
	return j.generate
}
func (j DeleteRepoJob) ExitChan() chan error {
	return j.exitChan
}
func (j DeleteRepoJob) Run() error {
	dbMgr := db.SatisDbManager{Path: j.dbPath}

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
func (j DeleteRepoJob) doDelete(repo string, repos []db.SatisRepository) ([]db.SatisRepository, error) {
	var err error = nil
	found := false

	rs := make([]db.SatisRepository, 0, len(repos))
	for _, r := range repos {
		if r.Url == repo {
			found = true
		} else {
			rs = append(rs, r)
		}
	}
	if !found {
		err = ErrRepoNotFound
	}
	return rs, err
}

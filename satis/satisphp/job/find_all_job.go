package job

import (
	"github.com/benschw/satis-go/satis/satisphp/db"
)

// Get all repos in the repo collection
func NewFindAllJob(dbPath string) *FindAllJob {
	return &FindAllJob{
		dbPath:    dbPath,
		exitChan:  make(chan error, 1),
		ReposResp: make(chan []db.SatisRepository, 1),
	}
}

type FindAllJob struct {
	ReposResp chan []db.SatisRepository
	dbPath    string
	generate  bool
	exitChan  chan error
}

func (j FindAllJob) Generate() bool {
	return j.generate
}
func (j FindAllJob) ExitChan() chan error {
	return j.exitChan
}
func (j FindAllJob) Run() error {
	dbMgr := db.SatisDbManager{Path: j.dbPath}

	err := dbMgr.Load()

	j.ReposResp <- dbMgr.Db.Repositories // might be empty

	return err // might be nil
}

package satisphp

import (
	"github.com/benschw/satis-go/satis/satisphp/db"
	"github.com/benschw/satis-go/satis/satisphp/job"
	"log"
)

var _ = log.Printf

type SatisJobProcessor struct {
	DbPath    string
	Jobs      chan job.SatisJob
	Generator Generator
}

// Run jobs added to Jobs chan
func (s *SatisJobProcessor) ProcessUpdates() {
	genCh := make(chan *db.SatisDbManager, 10)
	genRepoCh := make(chan string, 10)
	genExit := make(chan error, 1)

	go s.processGenerateJobs(genCh, genExit)
	go s.processGenerateRepoJobs(genRepoCh, genExit)

	for {
		j := <-s.Jobs
		err := j.Run()

		switch currentJob := j.(type) {
		// Generate Static Web
		case *job.GenerateJob:
			dbMgr := db.SatisDbManager{Path: s.DbPath}

			if err = dbMgr.Load(); err == nil {
				genCh <- &dbMgr
			}

		// Generate repo Static Web
		case *job.GenerateRepoJob:
			genRepoCh <- currentJob.RepoUrl()

		// Exit the generation goroutine
		case *job.ExitJob:
			genCh <- nil
			<-genExit
		}

		j.ExitChan() <- err

		// Stop Processing
		switch j.(type) {
		case *job.ExitJob:
			return
		}

	}
}

func (s *SatisJobProcessor) processGenerateJobs(genCh chan *db.SatisDbManager, genExit chan error) {
	for {
		dbMgr := <-genCh

		// Exit if mgr is nil
		if dbMgr == nil {
			genExit <- nil
			return
		}

		dbMgr.WriteStaging()

		// Do Static Site Generation
		if err := s.Generator.Generate(); err != nil {
			log.Print(err)
		}
	}
}


func (s *SatisJobProcessor) processGenerateRepoJobs(genPackageCh chan string, genExit chan error) {
	for {
		repoUrl := <-genPackageCh

		// Do Static Site Package Generation
		if err := s.Generator.GenerateRepo(repoUrl); err != nil {
			log.Print(err)
		}
	}
}

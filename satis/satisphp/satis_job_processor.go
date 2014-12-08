package satisphp

import (
	"github.com/benschw/satis-go/satis/satisphp/job"
	"log"
)

var _ = log.Printf

type SatisJobProcessor struct {
	Jobs      chan job.SatisJob
	Generator Generator
}

// Run jobs added to Jobs chan
func (s *SatisJobProcessor) ProcessUpdates() {
	genCh := make(chan job.SatisJob, 10)
	genExit := make(chan error, 1)

	go s.processGenerateJobs(genCh, genExit)

	for {
		j := <-s.Jobs
		err := j.Run()
		if err == nil && j.Generate() {
			genCh <- j
		}

		// Exit the generation goroutine
		switch j.(type) {
		case *job.ExitJob:
			genCh <- j
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

func (s *SatisJobProcessor) processGenerateJobs(genCh chan job.SatisJob, genExit chan error) {
	for {
		j := <-genCh

		// Exit if job type is `ExitJob`
		switch j.(type) {
		case *job.ExitJob:
			genExit <- nil
			return
		}

		// Do Static Site Generation
		if err := s.Generator.Generate(); err != nil {
			log.Print(err)
		}
	}
}

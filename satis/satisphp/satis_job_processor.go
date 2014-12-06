package satisphp

import (
	"log"
)

var _ = log.Printf

type SatisJobProcessor struct {
	Jobs      chan SatisJob
	Generator Generator
}

// TODO: have this just run "Run()" and move functionality to jobs
func (s *SatisJobProcessor) ProcessUpdates() {
	for {
		job := <-s.Jobs
		err := job.Run()

		if err == nil && job.Generate() {
			err = s.Generator.Generate()
		}

		job.ExitChan() <- err

		switch t := interface{}(job); t.(type) {
		case ExitJob:
			return
		}
	}

}

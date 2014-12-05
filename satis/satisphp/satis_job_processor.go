package satisphp

import (
	"log"
)

var _ = log.Printf

type SatisJobProcessor struct {
	DbPath    string
	Jobs      chan SatisJob
	Generator Generator
}

// TODO: have this just run "Run()" and move functionality to jobs
func (s *SatisJobProcessor) ProcessUpdates() {
	var err error = nil
	for {
		job := <-s.Jobs

		switch t := interface{}(job); t.(type) {
		// switch j := job.(type) {
		case SaveRepoJob:

			j := job.(SaveRepoJob)
			err = saveRepoInConfig(s.DbPath, j.repository)

		case FindAllJob:

			j := job.(FindAllJob)
			var repos []SatisRepository
			repos, err = findAllRepos(s.DbPath)
			j.reposResp <- repos // might be empty
		}

		if err == nil && job.Generate() {
			err = s.Generator.Generate()
		}

		job.ExitChan() <- err

		// if job.(type) == ExitJob && job.(ExitJob).exit {
		// 	return
		// }
	}

}

func saveRepoInConfig(dbPath string, repo SatisRepository) error {
	cfgMgr := SatisDbManager{Path: dbPath}

	if err := cfgMgr.Load(); err != nil {
		return err
	}

	if err := cfgMgr.SaveRepo(repo); err != nil {
		return err
	}

	if err := cfgMgr.Write(); err != nil {
		return err
	}
	return nil
}

func findAllRepos(dbPath string) ([]SatisRepository, error) {
	var repos []SatisRepository

	cfgMgr := SatisDbManager{Path: dbPath}

	if err := cfgMgr.Load(); err != nil {
		return repos, err
	}
	return cfgMgr.Db.Repositories, nil
}

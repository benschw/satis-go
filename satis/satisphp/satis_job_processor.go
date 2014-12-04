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

func (s *SatisJobProcessor) ProcessUpdates() {
	var err error = nil
	for {
		update := <-s.Jobs
		if update.Repository.Url != "" {
			err = saveRepoInConfig(s.DbPath, update.Repository)
		}
		if err == nil && update.Generate {
			err = s.Generator.Generate()
		}

		update.ExitChan <- err

		if update.Exit {
			return
		}
	}

}

func saveRepoInConfig(cfgPath string, repo SatisRepository) error {
	cfgMgr := SatisDbManager{Path: cfgPath}

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

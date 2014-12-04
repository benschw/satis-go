package satisapi

import (
	"log"
)

type SatisUpdateProcessor struct {
	ConfigPath string
	UpdateJobs chan UpdateJob
}

func (s *SatisUpdateProcessor) ProcessUpdates() {
	var err error = nil
	for {
		update := <-s.UpdateJobs
		if update.Repository.Url != "" {
			if err := saveRepoInConfig(s.ConfigPath, update); err != nil {
				log.Print(err)
			}
		}

		update.ExitChan <- err
		if update.Exit {
			return
		}
	}

}

func saveRepoInConfig(cfgPath string, update UpdateJob) error {
	cfgMgr := SatisConfigManager{Path: cfgPath}

	if err := cfgMgr.loadConfig(); err != nil {
		return err
	}

	if err := cfgMgr.saveRepo(update.Repository); err != nil {
		return err
	}

	if err := cfgMgr.writeConfig(); err != nil {
		return err
	}
	return nil
}

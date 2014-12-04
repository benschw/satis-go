package satisapi

import (
	"log"
)

type SatisUpdateProcessor struct {
	ConfigPath string
	UpdateJobs chan UpdateJob
}

func (s *SatisUpdateProcessor) ProcessUpdates() {
	for {
		update := <-s.UpdateJobs
		if err := processUpdate(s.ConfigPath, update); err != nil {
			log.Print(err)
		}
	}

}

func processUpdate(cfgPath string, update UpdateJob) error {
	if update.Repository.Url != "" {
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
	}
	return nil
}

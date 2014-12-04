package satisapi

import ()

type SatisManager struct {
	UpdateJobs chan UpdateJob
}

func (s *SatisManager) saveRepo(repo SatisRepository) error {
	job := UpdateJob{
		Repository: repo,
		ExitChan:   make(chan error),
	}

	s.UpdateJobs <- job

	return <-job.ExitChan
}

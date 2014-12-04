package satisapi

import ()

type SatisManager struct {
	UpdateJobs chan UpdateJob
}

func (s *SatisManager) saveRepo(repo SatisRepository) error {
	job := NewSaveRepoJob(repo, true, make(chan error))

	s.UpdateJobs <- *job

	return <-job.ExitChan
}

package satisphp

import ()

type SatisClient struct {
	Jobs chan SatisJob
}

func (s *SatisClient) SaveRepo(repo SatisRepository) error {
	job := NewSaveRepoJob(repo, true, make(chan error))

	return s.performJob(job)
}

func (s *SatisClient) GenerateSatisWeb() error {
	job := NewGenerateJob(make(chan error))

	return s.performJob(job)
}

func (s *SatisClient) performJob(job *SatisJob) error {
	s.Jobs <- *job

	return <-job.ExitChan
}

package satisphp

import (
	"log"
)

var _ = log.Print

type SatisClient struct {
	Jobs   chan SatisJob
	DbPath string
}

func (s *SatisClient) SaveRepo(repo SatisRepository) error {
	job := NewSaveRepoJob(s.DbPath, repo, true)

	return s.performJob(job)
}

func (s *SatisClient) DeleteRepo(repo string) error {
	job := NewDeleteRepoJob(s.DbPath, repo, true)

	return s.performJob(job)
}

func (s *SatisClient) FindAllRepos() ([]SatisRepository, error) {
	job := NewFindAllJob(s.DbPath)

	err := s.performJob(job)

	repos := <-job.reposResp
	return repos, err
}

func (s *SatisClient) GenerateSatisWeb() error {
	job := NewGenerateJob()

	return s.performJob(job)
}

func (s *SatisClient) performJob(job SatisJob) error {
	s.Jobs <- job

	return <-job.ExitChan()
}

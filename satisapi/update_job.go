package satisapi

import ()

type UpdateJob struct {
	Generate   bool
	Repository SatisRepository
	ExitChan   chan error
	Exit       bool
}

func NewGenerateJob(ch chan error) *UpdateJob {
	return &UpdateJob{
		Generate: true,
		ExitChan: ch,
		Exit:     false,
	}
}

func NewSaveRepoJob(repo SatisRepository, gen bool, ch chan error) *UpdateJob {
	return &UpdateJob{
		Generate:   gen,
		Repository: repo,
		ExitChan:   ch,
		Exit:       false,
	}
}

func NewExitJob(ch chan error) *UpdateJob {
	return &UpdateJob{
		ExitChan: ch,
		Exit:     true,
	}
}

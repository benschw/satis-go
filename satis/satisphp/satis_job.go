package satisphp

type SatisJob struct {
	Generate   bool
	Repository SatisRepository
	ExitChan   chan error
	Exit       bool
}

func NewGenerateJob(ch chan error) *SatisJob {
	return &SatisJob{
		Generate: true,
		ExitChan: ch,
		Exit:     false,
	}
}

func NewSaveRepoJob(repo SatisRepository, gen bool, ch chan error) *SatisJob {
	return &SatisJob{
		Generate:   gen,
		Repository: repo,
		ExitChan:   ch,
		Exit:       false,
	}
}

func NewExitJob(ch chan error) *SatisJob {
	return &SatisJob{
		ExitChan: ch,
		Exit:     true,
	}
}

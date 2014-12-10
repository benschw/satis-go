package job

// No op Job to signal job processor to exit
func NewExitJob() *ExitJob {
	return &ExitJob{
		exitChan: make(chan error, 1),
	}
}

type ExitJob struct {
	exitChan chan error
}

func (j ExitJob) ExitChan() chan error {
	return j.exitChan
}
func (j ExitJob) Run() error {
	return nil
}

package job

// No op job to signal job processor to regenerate satic web
func NewGenerateJob() *GenerateJob {
	return &GenerateJob{
		exitChan: make(chan error, 1),
	}
}

type GenerateJob struct {
	exitChan chan error
}

func (j GenerateJob) ExitChan() chan error {
	return j.exitChan
}
func (j GenerateJob) Run() error {
	return nil
}

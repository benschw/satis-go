package job

// No op job to signal job processor to regenerate one repo satic web
func NewGenerateRepoJob(repoUrl string) *GenerateRepoJob {
	return &GenerateRepoJob{
		repoUrl:  repoUrl,
		exitChan: make(chan error, 1),
	}
}

type GenerateRepoJob struct {
	repoUrl  string
	exitChan chan error
}

func (j GenerateRepoJob) ExitChan() chan error {
	return j.exitChan
}
func (j GenerateRepoJob) Run() error {
	return nil
}
func (j GenerateRepoJob) RepoUrl() string  {
	return j.repoUrl
}

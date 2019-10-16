package job

// No op job to signal job processor to regenerate one repo satic web
func NewGenerateRepoJob(packageName string) *GenerateRepoJob {
	return &GenerateRepoJob{
		packageName: packageName,
		exitChan:    make(chan error, 1),
	}
}

type GenerateRepoJob struct {
	packageName string
	exitChan    chan error
}

func (j GenerateRepoJob) ExitChan() chan error {
	return j.exitChan
}
func (j GenerateRepoJob) Run() error {
	return nil
}
func (j GenerateRepoJob) PackageName() string  {
	return j.packageName
}

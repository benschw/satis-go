package satisphp

type SatisJob interface {
	Generate() bool
	ExitChan() chan error
	//	Run() error
}

func NewGenerateJob() *GenerateJob {
	return &GenerateJob{
		exitChan: make(chan error, 1),
	}
}

type GenerateJob struct {
	exitChan chan error
}

func (j *GenerateJob) Generate() bool {
	return true
}
func (j *GenerateJob) ExitChan() chan error {
	return j.exitChan
}

func NewSaveRepoJob(repo SatisRepository, gen bool) *SaveRepoJob {
	return &SaveRepoJob{
		generate:   gen,
		repository: repo,
		exitChan:   make(chan error, 1),
	}
}

type SaveRepoJob struct {
	repository SatisRepository
	generate   bool
	exitChan   chan error
}

func (j *SaveRepoJob) Generate() bool {
	return j.generate
}
func (j *SaveRepoJob) ExitChan() chan error {
	return j.exitChan
}

func NewFindAllJob() *FindAllJob {
	return &FindAllJob{
		exitChan:  make(chan error, 1),
		reposResp: make(chan []SatisRepository, 1),
	}
}

type FindAllJob struct {
	reposResp chan []SatisRepository
	generate  bool
	exitChan  chan error
}

func (j *FindAllJob) Generate() bool {
	return j.generate
}
func (j *FindAllJob) ExitChan() chan error {
	return j.exitChan
}

func NewExitJob() *ExitJob {
	return &ExitJob{
		exitChan: make(chan error, 1),
	}
}

type ExitJob struct {
	exitChan chan error
}

func (j *ExitJob) Generate() bool {
	return false
}
func (j *ExitJob) ExitChan() chan error {
	return j.exitChan
}

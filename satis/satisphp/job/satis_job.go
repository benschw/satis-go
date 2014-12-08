package job

import (
	"errors"
)

var ErrRepoNotFound = errors.New("Repository Not Found")

type SatisJob interface {
	Generate() bool
	ExitChan() chan error
	Run() error
}

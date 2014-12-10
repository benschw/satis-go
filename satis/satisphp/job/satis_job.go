package job

import (
	"errors"
)

var ErrRepoNotFound = errors.New("Repository Not Found")

type SatisJob interface {
	ExitChan() chan error
	Run() error
}

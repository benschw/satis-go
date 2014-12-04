package satisapi

import ()

type UpdateJob struct {
	Repository SatisRepository
	ExitChan   chan error
}

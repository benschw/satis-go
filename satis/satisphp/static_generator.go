package satisphp

import (
	"log"
	"os/exec"
)

var _ = log.Print

type Generator interface {
	Generate() error
}

type StaticWebGenerator struct {
	DbPath    string
	SatisPath string
	WebPath   string
}

func (s *StaticWebGenerator) Generate() error {
	_, err := exec.
		Command(s.SatisPath+"bin/satis", "build", "--no-interaction", s.DbPath, s.WebPath).
		Output()

	return err
}

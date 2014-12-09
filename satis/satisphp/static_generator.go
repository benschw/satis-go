package satisphp

import (
	"github.com/benschw/satis-go/satis/satisphp/db"
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
		Command(s.SatisPath+"/bin/satis", "build", "--no-interaction", s.DbPath+db.StagingFile, s.WebPath).
		Output()

	return err
}

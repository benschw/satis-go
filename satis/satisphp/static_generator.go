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
	log.Print("Generating...")
	_, err := exec.
		Command(s.SatisPath+"/bin/satis", "build", s.DbPath+db.StagingFile, s.WebPath).
		Output()

	return err
}

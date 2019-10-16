package satisphp

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/benschw/satis-go/satis/satisphp/db"
)

var _ = log.Print

type Generator interface {
	Generate() error
	GenerateRepo(string) error
}

type StaticWebGenerator struct {
	DbPath  string
	WebPath string
}

func (s *StaticWebGenerator) Generate() error {
	log.Print("Generating...")
	out, err := exec.
		Command("satis", "--no-interaction", "build", s.DbPath+db.StagingFile, s.WebPath).
		CombinedOutput()
	if err != nil {
		log.Printf("Satis Generation Error: %s", string(out[:]))
	}
	return err
}

func (s *StaticWebGenerator) GenerateRepo(repoUrl string) error {
	log.Print(fmt.Sprintf(`Generating "%s" package...`, repoUrl))
	out, err := exec.
		Command("satis", "--no-interaction", "build", "--repository-url", repoUrl, s.DbPath+db.StagingFile, s.WebPath).
		CombinedOutput()
	if err != nil {
		log.Printf("Satis Generation Error: %s", string(out[:]))
	}
	return err
}

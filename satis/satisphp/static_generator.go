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
	GeneratePackage(string) error
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

func (s *StaticWebGenerator) GeneratePackage(repoPackage string) error {
	log.Print(fmt.Sprintf(`Generating "%s" package...`, repoPackage))
	out, err := exec.
		Command("satis", "--no-interaction", "build", s.DbPath+db.StagingFile, s.WebPath, repoPackage).
		CombinedOutput()
	if err != nil {
		log.Printf("Satis Generation Error: %s", string(out[:]))
	}
	return err
}

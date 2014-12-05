package main

import (
	"github.com/benschw/satis-go/satis"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	s := &satis.Server{
		DbPath:    dir + "/db.json",
		WebPath:   dir + "/web/",
		SatisPath: dir + "/lib/satis/",
		Bind:      ":8080",
		Name:      "My Repo",
		Homepage:  "http://localhost:8080",
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

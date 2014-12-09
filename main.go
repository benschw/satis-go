package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/benschw/satis-go/satis"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Dbpath    string
	Bind      string
	Satispath string
	Webpath   string
	Reponame  string
	Repohost  string
}

func getConfig(path string) (Config, error) {
	config := Config{}

	if _, err := os.Stat(path); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}

func main() {
	// Get Arguments
	var cfgPath string

	flag.StringVar(&cfgPath, "config", "/opt/satis/config.yaml", "Path to Config File")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <command> \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Get Command/Operation
	if flag.NArg() == 0 {
		flag.Usage()
		log.Fatal("Command argument required")
	}
	cmd := flag.Arg(0)

	// Load Config
	cfg, err := getConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Make Data Dir
	if err := os.MkdirAll(cfg.Dbpath, 0744); err != nil {
		log.Fatalf("Unable to create path: %v", err)
	}

	switch cmd {
	case "serve":
		// Configure Server
		s := &satis.Server{
			DbPath:    cfg.Dbpath,
			WebPath:   cfg.Webpath,
			SatisPath: cfg.Satispath,
			Bind:      cfg.Bind,
			Name:      cfg.Reponame,
			Homepage:  cfg.Repohost,
		}

		// Start Server
		if err := s.Run(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Unknown Command: %s", cmd)
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <command> \n", os.Args[0])
		flag.Usage()
	}

}

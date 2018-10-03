package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ghodss/yaml"
)

// A Tender provides attachment of a handler to http.
type Tender interface {
	DeployTender(path, route string) error
}

var configFile = flag.String("config-file", "config.yml", "Path to config file")

func main() {
	flag.Parse()

	log.Println("Loading config")
	b, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	var cfg config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on port", cfg.Port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}

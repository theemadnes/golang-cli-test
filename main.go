package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Revisions struct {
	Revision []struct {
		RevisionId string `yaml:"revision_id"`
		Region     string `yaml:"region"`
		ProjectId  string `yaml:"project_id"`
	} `yaml:"revisions"`
}

func readRevisions(filename string) (Revisions, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	//r := &revisions{}
	r := Revisions{}
	err = yaml.Unmarshal(buf, &r)
	if err != nil {
		return r, fmt.Errorf("in file %q: %v", filename, err)
	}

	return r, nil
}

func main() {

	// get file name from command line
	argsWithoutProg := os.Args[1:]

	fmt.Printf("Reading manifest file %v\n", string(argsWithoutProg[0]))

	r, err := readRevisions(argsWithoutProg[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", r.Revision[0].RevisionId)
	fmt.Printf("%v\n", r.Revision[1].RevisionId)

}

package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type resolverFile struct {
	location    string
	application string
	profile     string
}

func newResolverFile(location, application, profile string) *resolverFile {
	return &resolverFile{
		location:    location,
		application: application,
		profile:     profile,
	}
}

func (d *resolverFile) decode() (string, map[string]interface{}, error) {
	name, absolute, err := d.getFileName()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	content, err := ioutil.ReadFile(absolute)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	source := map[string]interface{}{}
	if err := yaml.Unmarshal(content, &source); err != nil {
		log.Println(err)
		return "", nil, err
	}

	return name, source, nil
}

func (d *resolverFile) getFileName() (string, string, error) {
	name := fmt.Sprintf("%s/%s-%s.yml", d.location, d.application, d.profile)

	content, err := filepath.Abs(name)
	if err != nil {
		return "", "", err
	}

	return name, content, nil
}

package service

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type resolveFile struct {
	location    string
	application string
	profile     string
}

func newResolveFile(location, application, profile string) *resolveFile {
	return &resolveFile{
		location:    location,
		application: application,
		profile:     profile,
	}
}

func (d *resolveFile) decode() (string, map[string]interface{}, error) {
	name, absolute, err := d.getFileName()
	if err != nil {
		logrus.Error(err)
		return "", nil, err
	}

	content, err := ioutil.ReadFile(absolute)
	if err != nil {
		logrus.Error(err)
		return "", nil, err
	}

	source := map[string]interface{}{}
	if err := yaml.Unmarshal(content, &source); err != nil {
		logrus.Error(err)
		return "", nil, err
	}

	return name, source, nil
}

func (d *resolveFile) getFileName() (string, string, error) {
	name := fmt.Sprintf("%s/%s-%s.yml", d.location, d.application, d.profile)

	content, err := filepath.Abs(name)
	if err != nil {
		return "", "", err
	}

	return name, content, nil
}

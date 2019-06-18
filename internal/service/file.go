package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/helderfarias/go-config-server/internal/domain"
	"gopkg.in/yaml.v2"
)

type fileDriveNative struct {
	searchLocations string
	application     string
	profile         string
}

func (e *fileDriveNative) BuildProperySources() []domain.PropertySource {
	name, absolute, err := e.MakeFileName()
	if err != nil {
		log.Println(err)
		return []domain.PropertySource{}
	}

	content, err := ioutil.ReadFile(absolute)
	if err != nil {
		log.Println(err)
		return []domain.PropertySource{}
	}

	source := map[string]interface{}{}
	if err := yaml.Unmarshal(content, &source); err != nil {
		log.Println(err)
		return []domain.PropertySource{}
	}

	return []domain.PropertySource{domain.PropertySource{
		Name:   name,
		Source: source,
	}}
}

func (e *fileDriveNative) MakeFileName() (string, string, error) {
	name := fmt.Sprintf("%s/%s-%s.yml", e.searchLocations, e.application, e.profile)

	content, err := filepath.Abs(name)
	if err != nil {
		return "", "", err
	}

	return name, content, nil
}

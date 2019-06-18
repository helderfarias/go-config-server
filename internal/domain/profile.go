package domain

import (
	"fmt"
	"path/filepath"

	"github.com/guregu/null"
)

type ProfileConfig struct {
	Name            null.String      `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           null.String      `json:"label"`
	Version         null.String      `json:"version"`
	State           null.String      `json:"state"`
	PropertySources []PropertySource `json:"propertySources"`
}

type PropertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}

type FileResolver struct {
	Name     string
	Absolute string
}

func (p *ProfileConfig) MakeFileName(searchLocation string) (FileResolver, error) {
	name := fmt.Sprintf("%s/%s-%s.yml", searchLocation, p.Name.String, p.Profiles[0])

	content, err := filepath.Abs(name)
	if err != nil {
		return FileResolver{}, err
	}

	return FileResolver{Name: name, Absolute: content}, nil
}

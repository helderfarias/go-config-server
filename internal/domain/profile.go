package domain

import (
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
	Index  int                    `json:"index"`
	Source map[string]interface{} `json:"source"`
}

type FileResolver struct {
	Name     string
	Absolute string
}

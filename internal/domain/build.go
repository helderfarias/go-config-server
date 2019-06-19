package domain

type BuildSource struct {
	Properties []PropertySource
	Options    map[string]string
}

func NewBuildSource() *BuildSource {
	return &BuildSource{
		Properties: []PropertySource{},
		Options:    map[string]string{},
	}
}

func (b *BuildSource) AddProperty(p PropertySource) *BuildSource {
	b.Properties = append(b.Properties, p)
	return b
}

func (b *BuildSource) AddOption(key, value string) *BuildSource {
	b.Options[key] = value
	return b
}

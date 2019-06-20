package service

import (
	"log"
	"strings"

	"github.com/helderfarias/go-config-server/internal/domain"
)

type fileDriveNative struct {
	source       map[string]interface{}
	application  string
	profile      string
	label        string
	cryptService CryptService
}

func (e *fileDriveNative) Build() *domain.BuildSource {
	directory := e.source["searchLocations"].(string)

	resolver := newResolverFile(directory, e.application, e.profile)

	name, data, err := resolver.decode()
	if err != nil {
		log.Println(err)
		return &domain.BuildSource{}
	}

	source := map[string]interface{}{}
	for key, value := range data {
		switch value.(type) {
		case string:
			source[key] = e.eval(value.(string))
		default:
			source[key] = value
		}
	}

	return domain.NewBuildSource().
		AddProperty(domain.PropertySource{
			Name:   name,
			Source: source,
		})
}

func (e *fileDriveNative) eval(source string) string {
	if strings.HasPrefix(source, "{cipher}") {
		content := strings.ReplaceAll(source, "{cipher}", "")
		content = strings.ReplaceAll(content, "\"", "")
		decoded, err := e.cryptService.Decrypt(content)
		if err != nil {
			log.Println(err)
			return source
		}
		return string(decoded)
	}

	return source
}

package service

import (
	"fmt"
	"strings"

	"github.com/hashicorp/vault/api"
	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/sirupsen/logrus"
)

type vaultDriveNative struct {
	source       map[string]interface{}
	application  string
	profile      string
	token        string
	label        string
	vaultToken   string
	index        int
	cryptService CryptService
}

func (e *vaultDriveNative) Build() *domain.BuildSource {
	client, err := api.NewClient(&api.Config{
		Address: e.source["uri"].(string),
	})

	if err != nil {
		logrus.Error(err)
		return domain.NewBuildSource()
	}

	client.SetToken(e.vaultToken)

	repo := fmt.Sprintf("secret/data/%s/%s", e.application, e.profile)
	secretValues, err := client.Logical().Read(repo)
	if err != nil {
		logrus.Error(err)
		return domain.NewBuildSource()
	}

	if secretValues == nil || secretValues.Data == nil{
		return domain.NewBuildSource()
	}

	source := map[string]interface{}{}
	if value := secretValues.Data["data"]; value != nil {
		source = value.(map[string]interface{})
	}

	metadata := map[string]interface{}{}
	if value := secretValues.Data["metadata"]; value != nil {
		metadata = value.(map[string]interface{})
	}

	name := e.source["uri"].(string)
	name = strings.Replace(name, "http", "vault", 1)
	name = strings.Replace(name, "https", "vault(s)", 1)

	return domain.NewBuildSource().
		AddOption("version", fmt.Sprintf("%v", metadata["version"])).
		AddProperty(domain.PropertySource{
			Name:   name,
			Source: source,
			Index:  e.index,
		})
}

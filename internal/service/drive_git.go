package service

import (
	"fmt"
	"strings"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type gitDriveNative struct {
	source       map[string]interface{}
	application  string
	profile      string
	label        string
	cryptService CryptService
}

func (e *gitDriveNative) Build() *domain.BuildSource {
	directory, repo, err := e.clone()
	if err != nil {
		logrus.Error(err)
		return domain.NewBuildSource()
	}

	resolver := newResolverFile(directory, e.application, e.profile)

	_, data, err := resolver.decode()
	if err != nil {
		logrus.Error(err)
		return domain.NewBuildSource()
	}

	head, err := repo.Head()
	if err != nil {
		logrus.Error(err)
		return domain.NewBuildSource()
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
		AddOption("version", fmt.Sprintf("%s", head.Hash())).
		AddProperty(domain.PropertySource{
			Name:   e.source["uri"].(string),
			Source: source,
		})
}

func (e *gitDriveNative) eval(source string) string {
	if strings.HasPrefix(source, "{cipher}") {
		content := strings.ReplaceAll(source, "{cipher}", "")
		content = strings.ReplaceAll(content, "\"", "")
		decoded, err := e.cryptService.Decrypt(content)
		if err != nil {
			logrus.Error(err)
			return source
		}
		return string(decoded)
	}

	return source
}

func (e *gitDriveNative) clone() (string, *git.Repository, error) {
	uri := e.source["uri"].(string)
	name := uri[strings.LastIndex(uri, "/")+1:]
	directory := fmt.Sprintf("/tmp/gcs/%s__%s", e.profile, name)

	cloneDir := e.source["clone_dir"]
	if cloneDir != nil && cloneDir.(string) != "" {
		directory = fmt.Sprintf("%s/%s__%s", cloneDir.(string), e.profile, name)
	}

	options := &git.CloneOptions{
		URL:           e.source["uri"].(string),
		ReferenceName: plumbing.NewBranchReferenceName(e.label),
	}

	if e.source["username"] != nil && e.source["password"] != nil {
		options.Auth = &http.BasicAuth{
			Username: e.source["username"].(string),
			Password: e.source["password"].(string),
		}
	}

	repo, err := git.PlainClone(directory, false, options)

	if err == git.ErrRepositoryAlreadyExists {
		repo, err := git.PlainOpen(directory)
		if err != nil {
			return directory, repo, err
		}

		repo, err = e.forcePullIf(repo)
		return directory, repo, err
	}

	repo, err = e.forcePullIf(repo)

	return directory, repo, err
}

func (e *gitDriveNative) forcePullIf(repo *git.Repository) (*git.Repository, error) {
	forcePull := e.source["force_pull"]
	if forcePull == nil || !forcePull.(bool) {
		return repo, nil
	}

	work, err := repo.Worktree()
	if err != nil {
		return repo, err
	}

	if err := work.Pull(&git.PullOptions{RemoteName: "origin"}); err == git.NoErrAlreadyUpToDate {
		return repo, nil
	}

	return repo, err
}

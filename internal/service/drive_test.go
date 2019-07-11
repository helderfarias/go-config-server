package service

import (
	"testing"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateDriveFileWhenActiveFile(t *testing.T) {
	c := domain.EnvConfig{}
	c.Cloud.Spring.Profiles.Active = "native "

	result := newDriveNative(c)

	assert.Equal(t, 1, len(result.(*composeDriveNative).targets))
	assert.IsType(t, &fileDriveNative{}, result.(*composeDriveNative).targets[0])
}

func TestShouldCreateDriveGitWhenActiveGit(t *testing.T) {
	c := domain.EnvConfig{}
	c.Cloud.Spring.Profiles.Active = " git"

	result := newDriveNative(c)

	assert.Equal(t, 1, len(result.(*composeDriveNative).targets))
	assert.IsType(t, &gitDriveNative{}, result.(*composeDriveNative).targets[0])
}

func TestShouldCreateDriveComposeWhenActives(t *testing.T) {
	c := domain.EnvConfig{}
	c.Cloud.Spring.Profiles.Active = " native, git "

	result := newDriveNative(c)

	assert.Equal(t, 2, len(result.(*composeDriveNative).targets))
	assert.IsType(t, &composeDriveNative{}, result)
	assert.IsType(t, &gitDriveNative{}, result.(*composeDriveNative).targets[0])
	assert.IsType(t, &fileDriveNative{}, result.(*composeDriveNative).targets[1])
}

func TestShouldEmptyDriveWhenEmptyActives(t *testing.T) {
	c := domain.EnvConfig{}
	c.Cloud.Spring.Profiles.Active = ""

	result := newDriveNative(c)

	assert.IsType(t, &emptyDriveNative{}, result)
}

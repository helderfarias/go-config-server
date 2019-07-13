package domain

type EnvConfig struct {
	Cloud       SpringCloudConfig
	Application string
	Profile     string
	Label       string
	VaultToken  string
}

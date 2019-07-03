package domain

type GithubCommit struct {
	Modified []string `json:"modified"`
}

type GithubEvent struct {
	Name    string         `json:"name"`
	Commits []GithubCommit `json:"commits"`
}

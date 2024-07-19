package github

import (
	"context"

	"github.com/google/go-github/v63/github"
	"github.com/wisdommatt/akpos-assignment/services/scm/providers"
)

type Provider struct {
	client *github.Client
}

func NewProvider(client *github.Client) *Provider {
	return &Provider{client: client}
}

func (p *Provider) Name() string {
	return "github"
}

func (p *Provider) ListRepositories(ctx context.Context, params providers.ListRepositoriesParams) ([]*providers.Repository, error) {
	// implementation
	return nil, nil
}

func (p *Provider) ListCommits(ctx context.Context, params providers.ListCommitsParams) ([]*providers.Commit, error) {
	// implementation
	return nil, nil
}

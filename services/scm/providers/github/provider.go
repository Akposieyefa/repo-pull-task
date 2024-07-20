package github

import (
	"context"
	"fmt"

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

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allRepos []*providers.Repository
	for {
		repos, resp, err := p.client.Repositories.List(ctx, "", opt)
		if err != nil {
			return nil, fmt.Errorf("error listing repositories: %w", err)
		}

		for _, repo := range repos {
			r := &providers.Repository{
				ID:   fmt.Sprintf("%d", repo.GetID()),
				Name: repo.GetName(),
			}
			allRepos = append(allRepos, r)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allRepos, nil
}

func (p *Provider) ListCommits(ctx context.Context, params providers.ListCommitsParams) ([]*providers.Commit, error) {

	opt := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allCommits []*providers.Commit
	for {
		commits, resp, err := p.client.Repositories.ListCommits(ctx, "Akposieyefa", "pedivel", opt)
		if err != nil {
			return nil, fmt.Errorf("error listing commits: %w", err)
		}

		for _, commit := range commits {
			c := &providers.Commit{
				ID:        commit.GetSHA(),
				Message:   commit.GetCommit().GetMessage(),
				Author:    commit.GetAuthor().GetLogin(),
				CreatedAt: commit.Author.CreatedAt.Time,
			}
			allCommits = append(allCommits, c)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allCommits, nil
}

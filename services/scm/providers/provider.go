//go:generate mockery --name Provider --outpkg testing --output ./testing --filename provider.go
package providers

import (
	"context"
	"time"
)

// Provider defines the interface for a source code management (SCM) provider.
// It encapsulates the basic methods required to interact with an SCM provider.
type Provider interface {
	// Name returns the name of the SCM provider.
	Name() string

	// ListRepositories returns a list of repositories based on the given parameters.
	// It takes a context and ListRepositoriesParams as input and returns a slice of Repository pointers or an error.
	ListRepositories(ctx context.Context, params ListRepositoriesParams) ([]*Repository, error)

	// ListCommits returns a list of commits for a given repository based on the given parameters.
	// It takes a context and ListCommitsParams as input and returns a slice of Commit pointers or an error.
	ListCommits(ctx context.Context, params ListCommitsParams) ([]*Commit, error)
}

// ListRepositoriesParams holds the parameters for listing repositories.
type ListRepositoriesParams struct {
	// Owner specifies the owner of the repositories.
	Owner string

	// Name specifies the name of the repository.
	Name string
}

// ListCommitsParams holds the parameters for listing commits in a repository.
type ListCommitsParams struct {
	// RepositoryOwner specifies the owner of the repository.
	RepositoryOwner string

	// RepositoryName specifies the name of the repository.
	RepositoryName string

	// AfterHash specifies the commit hash after which commits should be listed.
	AfterHash string

	// Limit specifies the maximum number of commits to be returned.
	Limit int

	// Since specifies the start time for the commits to be listed.
	Since time.Time

	// Until specifies the end time for the commits to be listed.
	Until time.Time
}

// Repository represents a source code repository.
type Repository struct {
	// ID is the unique identifier of the repository.
	ID string

	// Owner is the owner of the repository.
	Owner string

	// Name is the name of the repository.
	Name string

	// URL is the URL to access the repository.
	URL string
}

// Commit represents a commit in a source code repository.
type Commit struct {
	// ID is the unique identifier of the commit.
	ID string

	// Message is the commit message.
	Message string

	// Author is the author of the commit.
	Author string

	// CreatedAt is the time when the commit was created.
	CreatedAt time.Time
}

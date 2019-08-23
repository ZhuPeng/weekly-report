package github

import (
	"context"

	v3 "github.com/google/go-github/github"
	v4 "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type client struct {
	token string
	ctx   context.Context
	v4Cli *v4.Client
	v3Cli *v3.Client
}

func NewClient() *client {
	gcli := v4.NewClient(nil)
	return &client{v4Cli: gcli}
}

func NewClientWithToken(token string) *client {
	ctx := context.Background()
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, src)

	gcli := v4.NewClient(httpClient)
	return &client{token: token, ctx: ctx, v4Cli: gcli, v3Cli: v3.NewClient(httpClient)}
}

type Meta struct {
	ForkCount  int
	Forks      struct{ TotalCount int }
	Stargazers struct{ TotalCount int }
	Watchers   struct{ TotalCount int }
}

func (c *client) GetMeta(owner, repo string) (Meta, error) {
	var q struct {
		Repository Meta `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`
	}
	variables := map[string]interface{}{
		"repositoryOwner": v4.String(owner),
		"repositoryName":  v4.String(repo),
	}
	err := c.v4Cli.Query(context.Background(), &q, variables)
	if err != nil {
		return Meta{}, err
	}
	return q.Repository, nil
}

func (c *client) GetPR(owner, repo, state string) ([]*v3.PullRequest, error) {
	opt := &v3.PullRequestListOptions{State: state}
	prs, _, err := c.v3Cli.PullRequests.List(c.ctx, owner, repo, opt)
	if err != nil {
		return prs, err
	}
	return prs, nil
}

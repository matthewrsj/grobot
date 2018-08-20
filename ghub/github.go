package ghub

import (
	"context"

	"github.com/google/go-github/github"
)

func GetAllRepositories(org string) ([]*github.Repository, error) {
	c := github.NewClient(nil)
	ctx := context.Background()

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := c.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return allRepos, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allRepos, nil
}

func GetRepositoryByName(org, repoName string) (*github.Repository, error) {
	c := github.NewClient(nil)
	ctx := context.Background()
	repo, _, err := c.Repositories.Get(ctx, org, repoName)
	return repo, err
}

func GetOpenIssues(org, repoName string) ([]*github.Issue, error) {
	c := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.IssueListByRepoOptions{
		State:       "all",
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allIssues []*github.Issue
	for {
		issues, resp, err := c.Issues.ListByRepo(ctx, org, repoName, opt)
		if err != nil {
			return allIssues, err
		}
		realIssues := []*github.Issue{}
		for i := range issues {
			if issues[i].PullRequestLinks != nil {
				continue
			}
			realIssues = append(realIssues, issues[i])
		}
		allIssues = append(allIssues, realIssues...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allIssues, nil
}

func GetNumOpenClosedIssues(org, repoName string) (open, closed int, err error) {
	is, err := GetOpenIssues(org, repoName)
	if err != nil {
		return
	}

	for i := range is {
		switch is[i].GetState() {
		case "open":
			open++
		default:
			closed++
		}
	}
	return
}

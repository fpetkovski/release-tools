package compiler

import (
	"fmt"

	"github.com/fpetkovski/release-tools/pkg/changelog"
	"github.com/fpetkovski/release-tools/pkg/client"
	"github.com/fpetkovski/release-tools/pkg/gitrepo"
	"github.com/google/go-github/v41/github"
)

type Compiler interface {
	CompileChangelog(analysis *gitrepo.ChangeSet) (*changelog.ChangeLog, error)
}

func NewPullGenerator() (*generator, error) {
	gh := github.NewClient(
		client.NewBasicAuthClient(
			"fpetkovski",
			"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",
		),
	)
	analyzer, err := newPRAnalyzer(gh, "prometheus-operator", "prometheus-operator")
	if err != nil {
		return nil, err
	}

	r, err := gitrepo.New("/Users/filippetkovski/Projects/prometheus-operator")
	if err != nil {
		return nil, err
	}

	return &generator{
		compiler: analyzer,
		repo:     r,
	}, nil
}

type generator struct {
	compiler Compiler
	repo     *gitrepo.Repo
}

func (g *generator) Run() error {
	analysis, err := g.repo.Analyze()
	if err != nil {
		return err
	}

	changeLog, err := g.compiler.CompileChangelog(analysis)
	if err != nil {
		return err
	}
	fmt.Println(changeLog.Write(changelog.GroupFormatter{}))

	return err
}

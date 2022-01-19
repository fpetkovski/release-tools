package compiler

import (
	"context"
	"regexp"
	"strings"

	"github.com/fpetkovski/release-tools/pkg/changelog"
	"github.com/fpetkovski/release-tools/pkg/gitrepo"
	"github.com/google/go-github/v41/github"
)

type pullDescriptionAnalyser struct {
	gh    *github.Client
	org   string
	repo  string
	regex *regexp.Regexp
}

func newPRAnalyzer(gh *github.Client, org string, repo string) (Compiler, error) {
	regex, err := regexp.Compile("```release-note(.*)```")
	if err != nil {
		return nil, err
	}

	return &pullDescriptionAnalyser{
		gh:    gh,
		org:   org,
		repo:  repo,
		regex: regex,
	}, nil
}

func (g *pullDescriptionAnalyser) CompileChangelog(analysis *gitrepo.ChangeSet) (*changelog.ChangeLog, error) {
	ctx := context.Background()
	changeLog := changelog.New()
	prs := make(map[int64]struct{})

	for _, commit := range analysis.LastCommits {
		pulls, response, err := g.gh.PullRequests.ListPullRequestsWithCommit(
			ctx,
			g.org,
			g.repo,
			commit.Hash.String(),
			nil,
		)
		if response != nil && response.StatusCode == 422 {
			continue
		}
		if err != nil {
			return nil, err
		}

		if len(pulls) == 0 {
			continue
		}

		pr := pulls[0]
		if _, ok := prs[*pr.ID]; ok {
			continue
		}

		prs[*pr.ID] = struct{}{}
		message := ExtractChangelogEntryMessage(*pulls[0].Body)
		changeType := ExtractChangelogEntryType(*pulls[0].Body)
		changeLog.AddEntry(changelog.NewEntry(changeType, message, *pr.URL))
	}

	return changeLog, nil
}

func ExtractChangelogEntryMessage(text string) string {
	regex := regexp.MustCompile("(?s)```release-note(.*)```")
	matches := regex.FindStringSubmatch(text)
	if len(matches) != 2 {
		return ""
	}

	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(strings.TrimSpace(matches[1]), " ")
}

func ExtractChangelogEntryType(text string) changelog.ChangeType {
	patterns := map[string]string{
		"[x] `CHANGE`":      "change",
		"[x] `FEATURE`":     "feature",
		"[x] `BUGFIX`":      "bugfix",
		"[x] `ENHANCEMENT`": "enhancement",
	}

	for pattern, entryType := range patterns {
		if strings.Contains(text, pattern) {
			return changelog.ChangeType(entryType)
		}
	}

	return changelog.ChangeTypeNone
}

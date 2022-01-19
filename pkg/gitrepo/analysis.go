package gitrepo

import "github.com/go-git/go-git/v5/plumbing/object"

// ChangeSet is the type which describes the changes
// made to a git repository since the last git tag.
type ChangeSet struct {
	// LatestTag is the latest git tag in the git repository.
	LatestTag   string
	// LatestCommits are the git commits since the latest git tag.
	LastCommits []*object.Commit
}

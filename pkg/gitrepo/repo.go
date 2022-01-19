package gitrepo

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type Repo struct {
	repository *git.Repository
}

func New(path string) (*Repo, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	if err := w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash("a6629868e63e42b18ccd8f8f4ba3ee2b5bcfcc41"),
	}); err != nil {
		return nil, err
	}

	return &Repo{
		repository: r,
	}, nil
}

func (g *Repo) Analyze() (*ChangeSet, error) {
	head, err := g.repository.Head()
	if err != nil {
		return nil, err
	}

	log, err := g.repository.Log(&git.LogOptions{
		From:  head.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, err
	}

	tags, err := g.tagIndex()
	if err != nil {
		return nil, err
	}

	var latestTag string
	var lastCommits []*object.Commit
	if err := log.ForEach(func(commit *object.Commit) error {
		tag, ok := tags[commit.Hash.String()]
		if ok {
			latestTag = tag
			return storer.ErrStop
		}

		lastCommits = append(lastCommits, commit)
		return nil
	}); err != nil {
		return nil, err
	}

	return &ChangeSet{
		LatestTag:   latestTag,
		LastCommits: lastCommits,
	}, nil
}

func (g *Repo) tagIndex() (map[string]string, error) {
	tags, err := g.repository.Tags()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	if err := tags.ForEach(func(reference *plumbing.Reference) error {
		if err != nil {
			return err
		}

		targetReference, err := g.repository.ResolveRevision(plumbing.Revision(reference.Hash().String()))
		if err != nil {
			return err
		}

		result[targetReference.String()] = reference.Name().Short()
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

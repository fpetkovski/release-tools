package gitrepo

import (
	"fmt"
	"testing"
)

func TestGetAllTags(t *testing.T) {
	repo, err := New("/Users/filippetkovski/Projects/prometheus-operator")
	if err != nil {
		t.Fatal(err)
	}

	a, err := repo.Analyze()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a.LatestTag)
}

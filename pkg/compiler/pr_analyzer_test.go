package compiler

import (
	"github.com/fpetkovski/release-tools/pkg/changelog"
	"testing"
)

func TestExtractChangelogEntryType(t *testing.T) {
	tcs := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name: "feature type",
			text: `
- [ ] ` + "`CHANGE`" + `(fix or feature that would cause existing functionality to not work as expected)
- [x] ` + "`FEATURE`" + `(non-breaking change which adds functionality)
- [ ] ` + "`BUGFIX`" + `(non-breaking change which fixes an issue)
- [ ] ` + "`ENHANCEMENT`" + `(non-breaking change which improves existing functionality)
- [ ] ` + "`NONE`" + `(if none of the other choices apply. Example, tooling, build system, CI, docs, etc.)
`,
			expected: "feature",
		},
	}

	for _, ts := range tcs {
		actual := ExtractChangelogEntryType(ts.text)
		if actual != changelog.ChangeType(ts.expected) {
			t.Fatalf("invalid entry type, got %s want %s", actual, ts.expected)
		}
	}
}

func TestExtactChangelogEntry(t *testing.T) {
	tcs := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "inline release note",
			text:     "this is some description```release-note some release note```",
			expected: "some release note",
		},
		{
			name:     "empty release note",
			text:     "```release-note```",
			expected: "",
		},
		{
			name:     "empty release note with description",
			text:     "this is some description```release-note```",
			expected: "",
		},
		{
			name:     "empty description",
			text:     "",
			expected: "",
		},
		{
			name: "multiline release note",
			text: "this is some description " +
				"```release-note" +
				`some release 
				note ` +
				" ```",
			expected: "some release note",
		},
	}
	for _, tc := range tcs {
		actual := ExtractChangelogEntryMessage(tc.text)
		if tc.expected != actual {
			t.Fatalf("invalid release note, got %s, want %s", actual, tc.expected)
		}
	}
}
